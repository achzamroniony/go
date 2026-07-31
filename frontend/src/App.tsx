import { useState, useEffect } from 'react';
import './App.css';
import { checkHealth, fetchHello } from './services/api';
import type { HelloData, APIResponse } from './services/api';

function App() {
  const [apiStatus, setApiStatus] = useState<'loading' | 'online' | 'offline'>('loading');
  const [loadingData, setLoadingData] = useState<boolean>(false);
  const [apiResponse, setApiResponse] = useState<APIResponse<HelloData> | null>(null);
  const [helloData, setHelloData] = useState<HelloData | null>(null);

  // Poll backend health status on mount
  useEffect(() => {
    const verifyHealth = async () => {
      try {
        const response = await checkHealth();
        if (response.success) {
          setApiStatus('online');
        } else {
          setApiStatus('offline');
        }
      } catch {
        setApiStatus('offline');
      }
    };

    verifyHealth();
    // Re-check every 10 seconds
    const interval = setInterval(verifyHealth, 10000);
    return () => clearInterval(interval);
  }, []);

  // Fetch hello data from Go API
  const handleConnectAPI = async () => {
    setLoadingData(true);
    try {
      const response = await fetchHello();
      setApiResponse(response);
      if (response.success && response.data) {
        setHelloData(response.data);
      } else {
        setHelloData(null);
      }
    } catch (err: any) {
      setApiResponse({
        success: false,
        message: err.message || 'Koneksi gagal',
      });
      setHelloData(null);
    } finally {
      setLoadingData(false);
    }
  };

  return (
    <div className="app-container">
      {/* Header */}
      <header className="header">
        <span className="badge">Fullstack Development Template</span>
        <h1 className="text-gradient">Go + Fiber & React</h1>
        <p>Struktur folder standar profesional untuk memulai pembelajaran fullstack web app Anda</p>
      </header>

      {/* Main Grid */}
      <main className="main-grid">
        
        {/* Left Card: Architecture Info */}
        <section className="glass-card">
          <h2 style={{ marginBottom: '1rem', borderBottom: '1px solid var(--border-glass)', paddingBottom: '0.5rem' }}>
            Panduan Pembelajaran
          </h2>
          <p style={{ color: 'var(--text-secondary)', marginBottom: '1.5rem', fontSize: '0.95rem' }}>
            Template ini memisahkan backend (Go + Fiber) dan frontend (React + Vite) untuk pemisahan tugas (Separation of Concerns) yang jelas.
          </p>

          <h3 style={{ fontSize: '1.1rem', color: 'var(--color-secondary)' }}>Materi Utama:</h3>
          <ul className="topic-list">
            <li className="topic-item">
              <span>🚀</span> <strong>Go & Fiber:</strong> Performa tinggi, routing cepat, dan API yang clean.
            </li>
            <li className="topic-item">
              <span>🌐</span> <strong>RESTful API:</strong> Format standard JSON Response untuk konsistensi.
            </li>
            <li className="topic-item">
              <span>⚛️</span> <strong>React + TypeScript:</strong> UI interaktif, Type-safety, dan bundling kilat dengan Vite.
            </li>
            <li className="topic-item">
              <span>🛡️</span> <strong>CORS & Security:</strong> Penanganan kebijakan keamanan antar-origin secara tepat.
            </li>
          </ul>
        </section>

        {/* Right Card: Interactive Sandbox API */}
        <section className="glass-card api-section">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <h2 style={{ fontSize: '1.5rem' }}>API Testing Console</h2>
            
            <div className="api-status">
              <span>Status API:</span>
              <span className={`status-dot ${apiStatus === 'online' ? 'online' : apiStatus === 'offline' ? 'offline' : ''}`}></span>
              <strong style={{ 
                color: apiStatus === 'online' ? 'var(--success)' : apiStatus === 'offline' ? 'var(--error)' : 'var(--text-muted)',
                fontSize: '0.85rem',
                textTransform: 'uppercase'
              }}>
                {apiStatus === 'online' ? 'ONLINE' : apiStatus === 'offline' ? 'OFFLINE' : 'CHECKING...'}
              </strong>
            </div>
          </div>

          <p style={{ color: 'var(--text-secondary)', fontSize: '0.95rem' }}>
            Klik tombol di bawah ini untuk mengirim request HTTP GET ke endpoint <code>/api/hello</code> di backend Go Fiber Anda.
          </p>

          <div>
            <button 
              className="btn btn-primary" 
              onClick={handleConnectAPI}
              disabled={loadingData}
            >
              {loadingData ? 'Menghubungkan...' : 'Hubungkan ke API Go'}
            </button>
          </div>

          {/* Show Data when available */}
          {helloData && (
            <div style={{ marginTop: '1rem', animation: 'fadeIn 0.5s ease-out' }}>
              <h3 style={{ color: 'var(--success)', marginBottom: '0.5rem', display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
                <span>✅</span> Berhasil Terhubung!
              </h3>
              <p style={{ fontWeight: '500', fontSize: '1.05rem', marginBottom: '1rem' }}>
                "{helloData.message}"
              </p>
              
              <h4 style={{ fontSize: '0.95rem', color: 'var(--text-secondary)', marginBottom: '0.5rem' }}>
                Topik yang akan Anda pelajari:
              </h4>
              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '0.5rem', marginBottom: '1.5rem' }}>
                {helloData.topics.map((topic, i) => (
                  <span key={i} className="badge" style={{ backgroundColor: 'rgba(6, 182, 212, 0.1)', color: '#22d3ee', borderColor: 'rgba(6, 182, 212, 0.2)' }}>
                    {topic}
                  </span>
                ))}
              </div>
              <span style={{ fontSize: '0.8rem', color: 'var(--text-muted)' }}>
                Versi API: {helloData.version}
              </span>
            </div>
          )}

          {/* Raw JSON Display */}
          {apiResponse && (
            <div style={{ marginTop: '1rem' }}>
              <h4 style={{ fontSize: '0.9rem', color: 'var(--text-secondary)', marginBottom: '0.5rem' }}>
                HTTP JSON Response Payload:
              </h4>
              <pre className="response-box">
                {JSON.stringify(apiResponse, null, 2)}
              </pre>
            </div>
          )}
        </section>
      </main>

      {/* Footer */}
      <footer className="footer">
        <p>Go Fiber + React Fullstack Template | Dirancang untuk kebutuhan pembelajaran profesional</p>
        <p style={{ marginTop: '0.5rem', fontSize: '0.75rem' }}>Silakan jalankan backend di port 8080 dan frontend di port 5173</p>
      </footer>
    </div>
  );
}

export default App;
