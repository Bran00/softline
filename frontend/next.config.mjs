/** @type {import('next').NextConfig} */

const myBack = process.env.BACKNODE

const nextConfig = {
  env: {
    BACKNODE: myBack,
  },
};

export default nextConfig;
