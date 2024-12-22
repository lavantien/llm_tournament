import Layout from '../components/Layout';

export default function Home() {
  const loadMockData = async () => {
    try {
      const response = await fetch('/api/scores', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ action: 'loadMockData' })
      });
      if (response.ok) {
        alert('Mock data loaded successfully');
      } else {
        alert('Failed to load mock data');
      }
    } catch (error) {
      console.error('Error loading mock data:', error);
    }
  };

  const wipeData = async () => {
    try {
      const response = await fetch('/api/scores', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ action: 'wipeData' })
      });
      if (response.ok) {
        alert('Data wiped successfully');
        // Send signal to components to empty their state
        const event = new Event('dataWiped');
        window.dispatchEvent(event);
        console.log('Data wiped event dispatched');
      } else {
        alert('Failed to wipe data');
      }
    } catch (error) {
      console.error('Error wiping data:', error);
    }
  };

  return (
    <Layout>
      <h1 className="text-3xl font-bold mb-4">Welcome to LLM Benchmark</h1>
      <p className="mb-4">Use the navigation menu to access different sections of the app.</p>
      <button className="bg-mystic-highlight text-mystic-primary px-4 py-2 rounded hover:bg-mystic-accent" onClick={loadMockData}>Load Mock Data into DB</button>
      <button className="bg-mystic-highlight text-mystic-primary px-4 py-2 rounded hover:bg-mystic-accent ml-4" onClick={wipeData}>Wipe All Data from DB</button>
    </Layout>
  );
}
