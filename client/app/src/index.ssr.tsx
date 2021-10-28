import React from 'react'
import ReactDOM from 'react-dom/server'
import './index.css'
import App from './App'

const appOutput = ReactDOM.renderToString(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)

// serverRenderer is the callback injected by the go runtime to pass the rendered application back.
serverRenderer.render(appOutput)
