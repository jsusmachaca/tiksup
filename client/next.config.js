/** @type {import('next').NextConfig} */

const nextConfig = {
  env: {
    GATEWAY_URL: process.env.GATEWAY_URL,
  },
  images: {
    domains: ['gqldqjebnxaxmslrgqwf.supabase.co'],
  },
};

module.exports = nextConfig;
