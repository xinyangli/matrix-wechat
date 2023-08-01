{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, ...}@inputs: {
    packages."x86_64-linux".default = let pkgs = import inputs.nixpkgs { system = "x86_64-linux"; }; in pkgs.buildGoModule rec {
      src = ./.;
      name = "matrix-wechat";
      vendorSha256 = "sha256-V+JPYCuO1dTgtZ3mttMabkkj3ftKDf8XJi6A/mWm4BY=";
      buildInputs = [
        pkgs.olm
      ];
    };
    devShells."x86_64-linux".default = let pkgs = import inputs.nixpkgs { system = "x86_64-linux"; }; in pkgs.mkShell {
      hardeningDisable = [ "fortify" ];
      buildInputs = [ pkgs.go_1_20 ];
      packages = with pkgs; [
        gopls
        delve
        go-tools
        go_1_20
      ];
    };
  };
}
