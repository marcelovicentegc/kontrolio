function get() {
  try {
    const getBinary = require("./getBinary");
    return getBinary();
  } catch (err) {}
}

const binary = get();

if (binary) {
  binary.uninstall();
}
