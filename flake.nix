{
  description = "Set up your gitignore and license files like using a lumberjack";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };

        version =
          let
            rev = self.shortRev or (self.dirtyRev or "dirty");
            date = self.lastModifiedDate or "19700101";
          in
          "0.0.0-dev.${rev}-${date}";

        downjack = pkgs.buildGoModule {
          pname = "downjack";
          inherit version;

          src = self;

          subPackages = [ "." ];

          enableCGO = false;

          nativeBuildInputs = [ pkgs.installShellFiles ];

          ldflags = [
            "-s"
            "-w"
            "-X"
            "go.chardoncs.dev/downjack/internal/version.Version=${version}"
          ];

          vendorHash = "sha256-ZFp69rRijF0YXgHjx2bTeC+drF7sifIXsbO+OdnMeYI=";

          doCheck = false;

          postInstall = ''
            installShellCompletion --bash --name downjack \
              <($out/bin/downjack completion bash)
            installShellCompletion --zsh --name _downjack \
              <($out/bin/downjack completion zsh)
            installShellCompletion --fish --name downjack \
              <($out/bin/downjack completion fish)
          '';

          meta = with pkgs.lib; {
            description = "Set up your gitignore and license files like using a lumberjack";
            homepage = "https://github.com/chardoncs/downjack";
            license = licenses.mit;
            mainProgram = "downjack";
            platforms = platforms.linux ++ platforms.darwin;
          };
        };
      in
      {
        packages.default = downjack;
        packages.downjack = downjack;

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            golangci-lint
          ];
        };
      });
}