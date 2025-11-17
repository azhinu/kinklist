import defaultKinkListData from '../default_kinklist.json';

export function getDefaultTemplate() {
  return {
    id: '',
    nickname: defaultKinkListData.nickname || '',
    ratings: defaultKinkListData.ratings || [],
    blocks: defaultKinkListData.blocks || [],
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString()
  };
}

