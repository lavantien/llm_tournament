import Link from 'next/link';

export default function Layout({ children }) {
  return (
    <div className="min-h-screen flex flex-col bg-mystic-primary text-mystic-text dark:bg-mystic-secondary dark:text-mystic-text">
      <header className="bg-mystic-accent p-4 shadow-md">
        <h1 className="text-2xl font-bold">LLM Benchmark</h1>
        <nav>
          <ul className="flex space-x-4">
            <li><Link href="/" className="hover:text-mystic-highlight">Home</Link></li>
            <li><Link href="/model-manager" className="hover:text-mystic-highlight">Model Manager</Link></li>
            <li><Link href="/profile-manager" className="hover:text-mystic-highlight">Profile Manager</Link></li>
            <li><Link href="/prompt-manager" className="hover:text-mystic-highlight">Prompt Manager</Link></li>
            <li><Link href="/leaderboard" className="hover:text-mystic-highlight">Leaderboard</Link></li>
          </ul>
        </nav>
      </header>
      <main className="flex-1 p-4">
        {children}
      </main>
      <footer className="bg-mystic-accent p-4 text-center shadow-md">
        <p>© lavantien</p>
      </footer>
    </div>
  );
}
