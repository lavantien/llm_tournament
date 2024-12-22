import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';
import { Form, Button, Container, ListGroup } from 'react-bootstrap';

export default function ProfileManager() {
  const [profiles, setProfiles] = useState([]);
  const [formData, setFormData] = useState({
    name: '',
    systemPrompt: '',
    dryMultiplier: '',
    dryBase: '',
    dryAllowedLength: '',
    dryPenaltyLastN: '',
    repeatPenalty: '',
    repeatLastN: '',
    topK: '',
    topP: '',
    minP: '',
    topA: '',
    xtcThreshold: '',
    xtcProbability: '',
    temperature: ''
  });
  const [selectedProfile, setSelectedProfile] = useState(null);

  useEffect(() => {
    // Fetch profiles from the API route
    const fetchProfiles = async () => {
      try {
        const response = await fetch('/api/profiles');
        const data = await response.json();
        setProfiles(data);
        console.log('Profiles set:', data);
      } catch (error) {
        console.error('Error fetching profiles:', error);
      }
    };

    fetchProfiles();

    // Listen for dataWiped event
    const handleDataWiped = () => {
      setProfiles([]);
      console.log('Data wiped, profiles reset');
    };

    window.addEventListener('dataWiped', handleDataWiped);

    return () => {
      window.removeEventListener('dataWiped', handleDataWiped);
    };
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
    console.log('Form data set:', { ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission (e.g., add profile to the database)
    setProfiles([...profiles, formData]);
    console.log('Profiles set:', [...profiles, formData]);
    setFormData({
      name: '',
      systemPrompt: '',
      dryMultiplier: '',
      dryBase: '',
      dryAllowedLength: '',
      dryPenaltyLastN: '',
      repeatPenalty: '',
      repeatLastN: '',
      topK: '',
      topP: '',
      minP: '',
      topA: '',
      xtcThreshold: '',
      xtcProbability: '',
      temperature: ''
    });
    console.log('Form data reset');
  };

  const handleProfileClick = (profile) => {
    setSelectedProfile(profile);
    console.log('Selected profile set:', profile);
  };

  const handleClosePopup = () => {
    setSelectedProfile(null);
    console.log('Selected profile reset');
  };

  return (
    <Layout>
      <Container>
        <h1 className="h3 mb-4">Profile Manager</h1>
        <Form onSubmit={handleSubmit} className="bg-secondary p-3 rounded shadow">
          <Form.Group controlId="formName">
            <Form.Label>Name:</Form.Label>
            <Form.Control type="text" name="name" value={formData.name} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formSystemPrompt">
            <Form.Label>System Prompt:</Form.Label>
            <Form.Control as="textarea" name="systemPrompt" value={formData.systemPrompt} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formDryMultiplier">
            <Form.Label>Dry Multiplier:</Form.Label>
            <Form.Control type="number" name="dryMultiplier" value={formData.dryMultiplier} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formDryBase">
            <Form.Label>Dry Base:</Form.Label>
            <Form.Control type="number" name="dryBase" value={formData.dryBase} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formDryAllowedLength">
            <Form.Label>Dry Allowed Length:</Form.Label>
            <Form.Control type="number" name="dryAllowedLength" value={formData.dryAllowedLength} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formDryPenaltyLastN">
            <Form.Label>Dry Penalty Last N:</Form.Label>
            <Form.Control type="number" name="dryPenaltyLastN" value={formData.dryPenaltyLastN} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formRepeatPenalty">
            <Form.Label>Repeat Penalty:</Form.Label>
            <Form.Control type="number" name="repeatPenalty" value={formData.repeatPenalty} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formRepeatLastN">
            <Form.Label>Repeat Last N:</Form.Label>
            <Form.Control type="number" name="repeatLastN" value={formData.repeatLastN} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formTopK">
            <Form.Label>Top K:</Form.Label>
            <Form.Control type="number" name="topK" value={formData.topK} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formTopP">
            <Form.Label>Top P:</Form.Label>
            <Form.Control type="number" name="topP" value={formData.topP} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formMinP">
            <Form.Label>Min P:</Form.Label>
            <Form.Control type="number" name="minP" value={formData.minP} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formTopA">
            <Form.Label>Top A:</Form.Label>
            <Form.Control type="number" name="topA" value={formData.topA} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formXtcThreshold">
            <Form.Label>XTC Threshold:</Form.Label>
            <Form.Control type="number" name="xtcThreshold" value={formData.xtcThreshold} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formXtcProbability">
            <Form.Label>XTC Probability:</Form.Label>
            <Form.Control type="number" name="xtcProbability" value={formData.xtcProbability} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formTemperature">
            <Form.Label>Temperature:</Form.Label>
            <Form.Control type="number" name="temperature" value={formData.temperature} onChange={handleChange} />
          </Form.Group>
          <Button type="submit">Add Profile</Button>
        </Form>
        <ListGroup>
          {profiles.map((profile, index) => (
            <ListGroup.Item key={index} action onClick={() => handleProfileClick(profile)}>
              {profile.name}
            </ListGroup.Item>
          ))}
        </ListGroup>
        <DetailPopup item={selectedProfile} onClose={handleClosePopup} />
      </Container>
    </Layout>
  );
}
