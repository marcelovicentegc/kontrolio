const { Binary } = require("binary-install");
const os = require("os");

function getPlatform() {
  const type = os.type();
  const arch = os.arch();

  if (type === "Linux" && arch === "x64") return "linux_amd64";
  if (type === "Darwin" && arch === "x64") return "darwin_amd64";

  throw new Error(`Unsupported platform: ${type} ${arch}`);
}

function getBinary() {
  const platform = getPlatform();
  const version = require("../package.json").version;
  const url = `https://github.com/ktrlio/kontrolio-cli/releases/download/v${version}/kontrolio-cli_${version}_${platform}.tar.gz`;
  const name = "kontrolio";
  return new Binary(name, url);
}

module.exports = getBinary;
