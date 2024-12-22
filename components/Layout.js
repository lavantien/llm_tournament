import Link from 'next/link';
import styles from '../styles/Layout.module.css';

export default function Layout({ children }) {
  return (
    <div>
      <header className={styles.header}>
        <h1 className={styles.title}>LLM Benchmark</h1>
        <nav>
          <ul className={styles.navList}>
            <li><Link href="/">Home</Link></li>
            <li><Link href="/model-manager">Model Manager</Link></li>
            <li><Link href="/profile-manager">Profile Manager</Link></li>
            <li><Link href="/prompt-manager">Prompt Manager</Link></li>
            <li><Link href="/leaderboard">Leaderboard</Link></li>
          </ul>
        </nav>
      </header>
      <main className={styles.main}>
        {children}
      </main>
      <footer className={styles.footer}>
        <p>© lavantien</p>
      </footer>
    </div>
  );
}
