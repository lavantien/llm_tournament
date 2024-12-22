import Link from 'next/link';

export default function Layout({ children }) {
  return (
    <div className="container">
      <header>
        <h1>LLM Benchmark</h1>
        <nav>
          <Link href="/">Home</Link>
          <Link href="/model-manager">Model Manager</Link>
          <Link href="/profile-manager">Profile Manager</Link>
          <Link href="/prompt-manager">Prompt Manager</Link>
          <Link href="/leaderboard">Leaderboard</Link>
        </nav>
      </header>
      <main>
        {children}
      </main>
      <footer>
        <p>&copy; lavantien</p>
      </footer>
    </div>
  );
}
