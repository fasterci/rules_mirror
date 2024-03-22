package mirror

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func ExecuteContext(ctx context.Context, fromLocation, to, digest string) error {
	ref, err := name.ParseReference(fromLocation)
	if err != nil {
		return err
	}
	logs.Debug.Printf("in: %s/%s:%s", ref.Context().RegistryStr(), ref.Context().RepositoryStr(), ref.Identifier())
	dstRef, err := name.ParseReference(to)
	if err != nil {
		return err
	}
	logs.Debug.Print("out:", dstRef)
	hash, err := v1.NewHash(digest)
	if err != nil {
		return err
	}
	ref, err = name.ParseReference(fmt.Sprintf("%s@%s", ref.Context(), hash.String()))
	if err != nil {
		return err
	}
	shadst := fmt.Sprintf("%s@%s", dstRef.Context(), hash.String())
	shaDstRef, err := name.ParseReference(shadst)
	if err != nil {
		return err
	}
	// check if dst exists already
	if _, ok := dstRef.(name.Tag); ok {
		logs.Progress.Printf("fetching manifest for %s", dstRef)
		dst, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
		if err == nil {
			// if the dst manifest exists, check if it's the same as the src
			logs.Progress.Printf("found manifest for %s", dstRef)
			if dst.Digest.String() == hash.String() {
				return nil
			}
		}
	} else {
		logs.Progress.Printf("fetching manifest for %s", shaDstRef)
		_, err = remote.Head(shaDstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
		if err == nil {
			// if the dst manifest exists, check if it's the same as the src
			logs.Progress.Printf("found manifest for %s", shaDstRef)
			return nil
		}

	}
	logs.Progress.Printf("fetching manifest for %s", ref)
	src, err := remote.Get(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		return err
	}
	if src.Digest != hash {
		return fmt.Errorf("src digest %s does not match expected %s", src.Digest, hash)
	}
	logs.Progress.Printf("fetching image for %s", ref)
	image, err := src.Image()
	if err != nil {
		return fmt.Errorf("unable to fetch source image %s: %v", ref, err)
	}
	logs.Progress.Printf("pushing image to %s", dstRef)
	if err := remote.Write(dstRef, image, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx)); err != nil {
		return err
	}

	return nil
}
