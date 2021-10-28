/**
 * Provides the global server renderer instance
 * that allows us to call a Go callback with the rendered
 * output of the app for server-side rendering.
 *
 * This MUST only be accessed by code running server-side.
 */
declare var serverRenderer: ServerRenderer

declare interface ServerRenderer {
  render(renderedApp: string): void
}

/**
 * Provides the preloaded state for a redux store,
 * this is to be used for both client-side and server-side
 * rendering.
 *
 * On the server-side the data fetching should be in the host
 * application (Go) and passed down into the v8go context for the
 * rendering on the server.
 */
declare var __PRELOADED_STATE__: AppState
