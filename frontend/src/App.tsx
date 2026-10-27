import { useEffect, useState } from 'react'

export default function App() {
  const [status, setStatus] = useState('checking...')

  useEffect(() => {
    fetch('/health')
      .then((res) => res.json())
      .then((data) => setStatus(data.status))
      .catch(() => setStatus('unreachable'))
  }, [])

  return (
    <main style={{ padding: 24, fontFamily: 'sans-serif' }}>
      <h1>Learn Platform</h1>
      <p>API status: {status}</p>
    </main>
  )
}