const API_BASE = '/api';

export async function getKinkList(id) {
  const response = await fetch(`${API_BASE}/kinklist/${id}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch kinklist: ${response.statusText}`);
  }
  return await response.json();
}

export async function saveKinkList(id, data) {
  const response = await fetch(`${API_BASE}/kinklist/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    throw new Error(`Failed to save kinklist: ${response.statusText}`);
  }
  return await response.json();
}

