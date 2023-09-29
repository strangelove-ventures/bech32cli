{
  description = "bech32cli";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs@{ flake-parts, gitignore, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        packages.default = pkgs.buildGoModule {
          vendorSha256 = "sha256-oS/87ih8C2ScQkz/iYOrjDiXJ7vN6CwT/GI8DjzeFNg=";
          name = "bech32cli";
          src = gitignore.lib.gitignoreSource ./.;
        };
      };
    };
}
