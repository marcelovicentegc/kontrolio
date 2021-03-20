const { Binary } = require("binary-install");
const os = require("os");
const cTable = require("console.table");

const error = (msg) => {
  console.error(msg);
  process.exit(1);
};

const { version, name, repository } = require("../package.json");

const supportedPlatforms = [
  {
    TYPE: "Linux",
    ARCHITECTURE: "x64",
    GOLANG_TARGET: "linux_amd64",
    BINARY_NAME: "kontrolio",
  },
  {
    TYPE: "Darwin",
    ARCHITECTURE: "x64",
    GOLANG_TARGET: "darwin_amd64",
    BINARY_NAME: "kontrolio",
  },
];

const getPlatformMetadata = () => {
  const type = os.type();
  const architecture = os.arch();

  for (let index in supportedPlatforms) {
    let supportedPlatform = supportedPlatforms[index];
    if (
      type === supportedPlatform.TYPE &&
      architecture === supportedPlatform.ARCHITECTURE
    ) {
      return supportedPlatform;
    }
  }

  error(
    `Platform with type "${type}" and architecture "${architecture}" is not supported by ${name}.\nYour system must be one of the following:\n\n${cTable.getTable(
      supportedPlatforms
    )}`
  );
};

const getBinary = () => {
  const { BINARY_NAME, GOLANG_TARGET } = getPlatformMetadata();
  const url = `${repository.url}/releases/download/v${version}/kontrolio-cli_${version}_${GOLANG_TARGET}.tar.gz`;
  return new Binary(BINARY_NAME, url);
};

const run = () => {
  try {
    const binary = getBinary();
    binary.run();
  } catch (err) {
    error(err.message);
  }
};

const install = () => {
  try {
    const binary = getBinary();
    binary.install();
  } catch (err) {
    error(err.message);
  }
};

module.exports = {
  install,
  run,
};
