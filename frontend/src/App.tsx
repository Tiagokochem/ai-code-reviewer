import { useState } from 'react'
import CodeReviewer from './components/CodeReviewer'
import './App.css'

function App() {
  return (
    <div className="app">
      <header className="app-header">
        <h1>🔥 AI Code Reviewer</h1>
        <p>Análise automática de código com IA</p>
      </header>
      <main className="app-main">
        <CodeReviewer />
      </main>
    </div>
  )
}

export default App
