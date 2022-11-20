package mirror

import (
	"context"
	"flag"
	"fmt"
	"path"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

var (
	FromLocation string
	ToPrefix     string
	Digest       string
)

func init() {
	flag.StringVar(&FromLocation, "from", "", "The location of the image to mirror, required")
	flag.StringVar(&ToPrefix, "to_prefix", "", "The registry and repository prefix to mirror the image to, required")
	flag.StringVar(&Digest, "digest", "", "The digest of the image, like sha256:1234, required")
}

func ExecuteContext(ctx context.Context) error {
	// verify that the flags are set
	if FromLocation == "" || ToPrefix == "" || Digest == "" {
		return flag.ErrHelp
	}
	ref, err := name.ParseReference(FromLocation)
	if err != nil {
		return err
	}
	logs.Debug.Printf("in: %s/%s:%s", ref.Context().RegistryStr(), ref.Context().RepositoryStr(), ref.Identifier())
	dst := fmt.Sprintf("%s:%s", path.Join(ToPrefix, ref.Context().RegistryStr(), ref.Context().RepositoryStr()), ref.Identifier())
	logs.Debug.Print("out:", dst)
	dstRef, err := name.ParseReference(dst)
	if err != nil {
		return err
	}
	hash, err := v1.NewHash(Digest)
	if err != nil {
		return err
	}
	shadst := fmt.Sprintf("%s@%s", path.Join(ToPrefix, ref.Context().RegistryStr(), ref.Context().RepositoryStr()), hash.String())
	shaDstRef, err := name.ParseReference(shadst)
	if err != nil {
		return err
	}
	// fetch dst manifest
	logs.Progress.Printf("fetching manifest for %s", shaDstRef)
	_, err = remote.Head(shaDstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err == nil {
		// if the dst manifest exists, check if it's the same as the src
		logs.Progress.Printf("found manifest for %s", shaDstRef)
		return nil
	}
	logs.Progress.Printf("fetching manifest for %s", ref)
	src, err := remote.Get(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return err
	}
	logs.Progress.Printf("fetching image for %s", ref)
	image, err := src.Image()
	if err != nil {
		return fmt.Errorf("unable to fetch source image %s: %v", ref, err)
	}
	imgDigest, err := image.Digest()
	if err != nil {
		return fmt.Errorf("unable to get digest for image %s: %v", ref, err)
	}
	if imgDigest != hash {
		return fmt.Errorf("src digest %s does not match expected %s", imgDigest, hash)
	}
	logs.Progress.Printf("pushing image to %s", dstRef)
	if err := remote.Write(dstRef, image, remote.WithAuthFromKeychain(authn.DefaultKeychain)); err != nil {
		return err
	}

	return nil
}
