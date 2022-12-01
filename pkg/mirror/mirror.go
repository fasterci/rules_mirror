package mirror

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

var (
	FromLocation string
	To           string
	Digest       string
)

func init() {
	flag.StringVar(&FromLocation, "from", "", "The location of the image to mirror, required")
	flag.StringVar(&To, "to", "", "The location of the mirror destination, required")
	flag.StringVar(&Digest, "digest", "", "The digest of the image, like sha256:1234, required")
}

func ExecuteContext(ctx context.Context) error {
	// verify that the flags are set
	if FromLocation == "" || To == "" || Digest == "" {
		return flag.ErrHelp
	}
	ref, err := name.ParseReference(FromLocation)
	if err != nil {
		return err
	}
	logs.Debug.Printf("in: %s/%s:%s", ref.Context().RegistryStr(), ref.Context().RepositoryStr(), ref.Identifier())
	dstRef, err := name.ParseReference(To)
	if err != nil {
		return err
	}
	logs.Debug.Print("out:", dstRef)
	hash, err := v1.NewHash(Digest)
	if err != nil {
		return err
	}
	shadst := fmt.Sprintf("%s@%s", dstRef.Context(), hash.String())
	shaDstRef, err := name.ParseReference(shadst)
	if err != nil {
		return err
	}
	// fetch dst manifest
	checkref := shaDstRef
	if _, ok := dstRef.(name.Tag); ok {
		checkref = dstRef
	}
	logs.Progress.Printf("fetching manifest for %s", checkref)
	_, err = remote.Head(checkref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err == nil {
		// if the dst manifest exists, check if it's the same as the src
		logs.Progress.Printf("found manifest for %s", checkref)
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
