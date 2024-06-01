# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "unstable";
  # Use https://search.nixos.org/packages to find packages
  packages = [
    pkgs.go_1_22
    pkgs.nodejs_20
    pkgs.nodePackages.nodemon
    pkgs.gnumake
    pkgs.docker-compose
  ];
  services = {
    docker = {
      enable = true;
    };
  };
  # Sets environment variables in the workspace
  env = {};
  idx = {
    # Search for the extensions you want on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      "golang.go"
    ];
     workspace = {
        onStart = {
        # Example: start a background task to watch and re-build backend code
        # watch-backend = "npm run watch-backend";
        dependencies = "go mod tidy";
      };
     };
  };
}
