import Link from 'next/link';

export default function Layout({ children }) {
  return (
    <div className="d-flex flex-column min-vh-100 bg-dark text-light">
      <header className="bg-secondary p-3 shadow-sm">
        <h1 className="h2">LLM Benchmark</h1>
        <nav>
          <ul className="nav">
            <li className="nav-item"><Link href="/" className="nav-link text-light">Home</Link></li>
            <li className="nav-item"><Link href="/model-manager" className="nav-link text-light">Model Manager</Link></li>
            <li className="nav-item"><Link href="/profile-manager" className="nav-link text-light">Profile Manager</Link></li>
            <li className="nav-item"><Link href="/prompt-manager" className="nav-link text-light">Prompt Manager</Link></li>
            <li className="nav-item"><Link href="/leaderboard" className="nav-link text-light">Leaderboard</Link></li>
          </ul>
        </nav>
      </header>
      <main className="flex-grow-1 p-3">
        {children}
      </main>
      <footer className="bg-secondary p-3 text-center shadow-sm">
        <p>&copy; lavantien</p>
      </footer>
    </div>
  );
}
