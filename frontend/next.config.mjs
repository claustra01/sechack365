class SkipLargeSourceCachePlugin {
  constructor({ threshold } = {}) {
    this.threshold = threshold ?? 128 * 1024;
  }

  apply(compiler) {
    compiler.hooks.compilation.tap('SkipLargeSourceCachePlugin', (compilation) => {
      compilation.hooks.succeedModule.tap('SkipLargeSourceCachePlugin', (module) => {
        if (!module?.resource || !module.resource.includes('node_modules')) {
          return;
        }
        const originalSource =
          typeof module.originalSource === 'function' ? module.originalSource() : null;
        if (!originalSource) {
          return;
        }

        let size = 0;
        if (typeof originalSource.size === 'function') {
          size = originalSource.size();
        } else {
          const value =
            typeof originalSource.source === 'function'
              ? originalSource.source()
              : undefined;
          if (typeof value === 'string') {
            size = Buffer.byteLength(value);
          } else if (value && typeof value.length === 'number') {
            size = value.length;
          }
        }

        if (size >= this.threshold) {
          module.buildInfo.cacheable = false;
        }
      });
    });
  }
}

/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  webpack(config) {
    config.plugins = config.plugins ?? [];
    config.plugins.push(new SkipLargeSourceCachePlugin());
    return config;
  },
};

export default nextConfig;
