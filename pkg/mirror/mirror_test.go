package mirror_test

import (
	"context"
	"testing"

	"github.com/fasterci/rules_mirror/pkg/mirror"
	"github.com/fasterci/rules_mirror/pkg/testing/testregistry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func TestExecuteContext_HappypathTag(t *testing.T) {
	r, cleanup := testregistry.SetupRegistry(t)
	defer cleanup()

	from := "gcr.io/distroless/base:nonroot-amd64"
	// note: this hash is not the very recent one, but it's valid actual hash of some previous version
	hash := "sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6435"
	to := r.Name() + "/distroless/base:nonroot-amd64"

	ctx := context.Background()
	if d, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	err := mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}

	dstRef, err := name.ParseReference(to)
	if err != nil {
		t.Fatalf("Failed to parse reference %s: %v", to, err)
	}
	r1, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r1.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r1.Digest)
	}

	err = mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}
	r2, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r2.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r2.Digest)
	}

}

func TestExecuteContext_UpgradeTag(t *testing.T) {
	r, cleanup := testregistry.SetupRegistry(t)
	defer cleanup()

	from := "gcr.io/distroless/base:nonroot-amd64"
	hash := "sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6435"
	to := r.Name() + "/distroless/base:nonroot-amd64"

	ctx := context.Background()
	if d, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	err := mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}

	dstRef, err := name.ParseReference(to)
	if err != nil {
		t.Fatalf("Failed to parse reference %s: %v", to, err)
	}
	r1, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r1.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r1.Digest)
	}

	hash2 := "sha256:ef4831f68d82eba8414e7f5eef8971810bf5b69a7a295216601fa84f7b2f4989"

	err = mirror.ExecuteContext(ctx, from, to, hash2)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}
	r2, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r2.Digest.String() != hash2 {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash2, r2.Digest)
	}

}

func TestExecuteContext_HappypathSha(t *testing.T) {
	r, cleanup := testregistry.SetupRegistry(t)
	defer cleanup()

	from := "gcr.io/distroless/base@sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6435"
	hash := "sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6435"
	to := r.Name() + "/distroless/base@sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6435"

	ctx := context.Background()
	if d, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	err := mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}

	dstRef, err := name.ParseReference(to)
	if err != nil {
		t.Fatalf("Failed to parse reference %s: %v", to, err)
	}
	r1, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r1.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r1.Digest)
	}

	err = mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}
	r2, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r2.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r2.Digest)
	}

}

func TestExecuteContext_HappypathIndexSha(t *testing.T) {
	r, cleanup := testregistry.SetupRegistry(t)
	defer cleanup()

	from := "gcr.io/distroless/base@sha256:1a8ece87bd75cde87d0484ef48eb60ea25811baf90967265956ae4fa2098dd9d"
	hash := "sha256:1a8ece87bd75cde87d0484ef48eb60ea25811baf90967265956ae4fa2098dd9d"
	to := r.Name() + "/distroless/base"

	ctx := context.Background()
	if d, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	err := mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}

	dstRef, err := name.ParseReference(to)
	if err != nil {
		t.Fatalf("Failed to parse reference %s: %v", to, err)
	}
	r1, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r1.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r1.Digest)
	}

	err = mirror.ExecuteContext(ctx, from, to, hash)
	if err != nil {
		t.Fatalf("Failed to mirror %s to %s: %v", from, to, err)
	}
	r2, err := remote.Get(dstRef, remote.WithAuthFromKeychain(authn.DefaultKeychain), remote.WithContext(ctx))
	if err != nil {
		t.Fatalf("Failed to fetch %s: %v", to, err)
	}
	if r2.Digest.String() != hash {
		t.Fatalf("Expected %s to have digest %s, got %s", to, hash, r2.Digest)
	}

}

func TestExecuteContext_BadSha(t *testing.T) {
	r, cleanup := testregistry.SetupRegistry(t)
	defer cleanup()

	from := "gcr.io/distroless/base:nonroot-amd64"
	// note: this hash is not the very recent one, but it's valid actual hash of some previous version
	hash := "sha256:1c9093af306ef03503b8450b08fe6a2a13ba6d2c697ff74031a915f9201f6434"
	to := r.Name() + "/distroless/base:nonroot-amd64"

	ctx := context.Background()
	if d, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	err := mirror.ExecuteContext(ctx, from, to, hash)
	if err == nil {
		t.Fatalf("error expected")
	}
	if err.Error() != "source image gcr.io/distroless/base:nonroot-amd64 has digest sha256:ef4831f68d82eba8414e7f5eef8971810bf5b69a7a295216601fa84f7b2f4989" {
		t.Fatalf("unexpected error %v", err)
	}

}
