import Link from 'next/link';
import { Navbar, Nav, Container } from 'react-bootstrap';

export default function Layout({ children }) {
  return (
    <Container fluid className="d-flex flex-column min-vh-100 bg-dark text-light">
      <header>
        <Navbar bg="secondary" variant="dark" expand="lg">
          <Container fluid>
            <Navbar.Brand href="#">LLM Benchmark</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="me-auto">
                <Link href="/" passHref legacyBehavior><Nav.Link>Home</Nav.Link></Link>
                <Link href="/model-manager" passHref legacyBehavior><Nav.Link>Model Manager</Nav.Link></Link>
                <Link href="/profile-manager" passHref legacyBehavior><Nav.Link>Profile Manager</Nav.Link></Link>
                <Link href="/prompt-manager" passHref legacyBehavior><Nav.Link>Prompt Manager</Nav.Link></Link>
                <Link href="/leaderboard" passHref legacyBehavior><Nav.Link>Leaderboard</Nav.Link></Link>
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
      </header>
      <main className="flex-grow-1 p-3">
        <Container>
          {children}
        </Container>
      </main>
      <footer className="bg-secondary text-center p-3">
        <Container>
          <p className="text-light">&copy; lavantien</p>
        </Container>
      </footer>
    </Container>
  );
}
