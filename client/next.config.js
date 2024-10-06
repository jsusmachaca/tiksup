/** @type {import('next').NextConfig} */

const nextConfig = {
  env: {
    NODE_URL: process.env.NODE_URL,
  },
};

module.exports = nextConfig;
