// ignition/modules/Counter.js
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("CounterModule", (m) => {
  // 如果构造函数需要参数，在这里传；本合约无参
  const counter = m.contract("Counter");

  return { counter };
});