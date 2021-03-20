#!/usr/bin/env node

const child = require("child_process");
const fs = require("fs");
const os = require("os");
const path = require("path");
const cTable = require("console.table");
const axios = require("axios");
const rimraf = require("rimraf");
const tar = require("tar");

const { version, repository } = require("./package.json");

const binDir = path.join(__dirname, "bin");
const appName = "kontrolio";

// Determine the install directory by version so that we can detect when we need
// to upgrade to a new version.
const installDir = path.join(binDir, version, appName);

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

const platform = (() => {
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
    `Platform with type "${type}" and architecture "${architecture}" is not supported by ${appName}.\nYour system must be one of the following:\n\n${cTable.getTable(
      supportedPlatforms
    )}`
  );
})();

const binName = platform.TYPE === "Windows" ? ".exe" : "";
const binPath = path.join(installDir, binName);

const install = async () => {
  const { GOLANG_TARGET } = platform;
  const url = `${repository.url}/releases/download/v${version}/kontrolio-cli_${version}_${GOLANG_TARGET}.tar.gz`;

  fs.mkdirSync(installDir, { recursive: true });

  try {
    const response = await axios({ url, responseType: "stream" });

    // Strip the outer directory when extracting. Just get the binary.
    const tarWriter = tar.extract({ strip: 1, cwd: installDir });
    response.data.pipe(tarWriter);

    // Need to return a promise with the writer to ensure we can await for it to complete.
    return new Promise((resolve, reject) => {
      tarWriter.on("finish", resolve);
      tarWriter.on("error", reject);
    });
  } catch (err) {
    throw new Error(`Download archive ${url}: ${err.message}`);
  }
};

const run = async () => {
  if (!fs.existsSync(binPath)) {
    // Remove any existing binaries before installing the new one.
    rimraf.sync(binDir);
    console.log(`Installing kontrolio ${version}...`);
    try {
      await install();
    } catch (err) {
      console.error(`Failed to install: ${err.message}`);
      process.exit(1);
    }
    console.log("Install completed.");
  }

  const [, , ...args] = process.argv;
  const options = { cwd: process.cwd(), stdio: "inherit" };
  const { status, error } = child.spawnSync(binPath, args, options);

  if (error) {
    console.error(error);
    process.exit(1);
  }

  process.exit(status);
};

run();
