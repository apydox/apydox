#!/bin/bash
# Bundles the application for server-side rendering with esbuild.

npx esbuild \
    src/index.ssr.tsx \
    --inject:src/react-shim.js \
    --bundle \
    --sourcemap \
    --outfile=build-ssr/out.js \
    --loader:.svg=text \
    --define:process.env.NODE_ENV=\"production\"
