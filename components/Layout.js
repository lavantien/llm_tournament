import Link from 'next/link';

export default function Layout({ children }) {
  return (
    <div>
      <header>
        <nav>
          <ul>
            <li><Link href="/">Home</Link></li>
            <li><Link href="/model-manager">Model Manager</Link></li>
            <li><Link href="/profile-manager">Profile Manager</Link></li>
            <li><Link href="/prompt-manager">Prompt Manager</Link></li>
            <li><Link href="/leaderboard">Leaderboard</Link></li>
          </ul>
        </nav>
      </header>
      <main>
        {children}
      </main>
      <footer>
        <p>&copy; 2023 Spreadsheet-like Web App</p>
      </footer>
    </div>
  );
}
