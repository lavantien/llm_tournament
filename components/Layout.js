import Link from 'next/link';

export default function Layout({ children }) {
  return (
    <div className="min-h-screen flex flex-col">
      <header className="bg-gray-800 text-white p-4">
        <h1 className="text-2xl font-bold">LLM Benchmark</h1>
        <nav>
          <ul className="flex space-x-4">
            <li><Link href="/">Home</Link></li>
            <li><Link href="/model-manager">Model Manager</Link></li>
            <li><Link href="/profile-manager">Profile Manager</Link></li>
            <li><Link href="/prompt-manager">Prompt Manager</Link></li>
            <li><Link href="/leaderboard">Leaderboard</Link></li>
          </ul>
        </nav>
      </header>
      <main className="flex-1 p-4">
        {children}
      </main>
      <footer className="bg-gray-800 text-white p-4 text-center">
        <p>© lavantien</p>
      </footer>
    </div>
  );
}
