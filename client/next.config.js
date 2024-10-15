/** @type {import('next').NextConfig} */

const nextConfig = {
  env: {
    GATEWAY_URL: process.env.GATEWAY_URL,
  },
};

module.exports = nextConfig;
