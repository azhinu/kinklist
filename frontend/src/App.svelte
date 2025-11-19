<script>
  import { onMount } from 'svelte';
  import { getKinkList, saveKinkList } from './lib/api.js';
  import { getDefaultTemplate } from './lib/defaultTemplate.js';
  import Block from './lib/Block.svelte';
  import EditRatingsModal from './lib/EditRatingsModal.svelte';
  import EditBlockModal from './lib/EditBlockModal.svelte';
  import ShareLinkModal from './lib/ShareLinkModal.svelte';
  import WelcomeModal from './lib/WelcomeModal.svelte';
  import './styles.css';

  let kinkList = null;
  let loading = true;
  let error = null;
  let saving = false;
  let showRatingsModal = false;
  let showBlockModal = false;
  let editingBlock = null;
  let editingBlockIndex = -1;
  let showShareModal = false;
  let shareLink = '';
  let showWelcomeModal = false;

  // Extract UUID from URL path
  function getUUIDFromPath() {
    const path = window.location.pathname;
    const match = path.match(/\/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})/i);
    return match ? match[1] : null;
  }

  // Generate UUID v4
  function generateUUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
      const r = Math.random() * 16 | 0;
      const v = c === 'x' ? r : (r & 0x3 | 0x8);
      return v.toString(16);
    });
  }

  onMount(async () => {
    // Check if this is the first visit
    if (typeof window !== 'undefined') {
      const hasSeenWelcome = localStorage.getItem('kinklist_welcome_seen');
      if (!hasSeenWelcome) {
        showWelcomeModal = true;
      }
    }

    try {
      const uuid = getUUIDFromPath();
      if (uuid) {
        kinkList = await getKinkList(uuid);
      } else {
        kinkList = getDefaultTemplate();
      }
    } catch (err) {
      error = err.message;
      // Fallback to default template on error
      kinkList = getDefaultTemplate();
    } finally {
      loading = false;
    }
  });

  function handleSave() {
    if (!kinkList) return;
    
    saving = true;
    // Always generate a new UUID for each save
    const uuid = generateUUID();
    
    // Ensure data structure matches Go model (snake_case)
    const dataToSave = {
      id: uuid,
      nickname: kinkList.nickname || '',
      ratings: kinkList.ratings || [],
      blocks: kinkList.blocks || [],
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    };

    saveKinkList(uuid, dataToSave)
      .then((saved) => {
        kinkList = saved;
        // Update URL without reload
        window.history.pushState({}, '', `/${uuid}`);
        // Generate share link and open modal
        shareLink = `${window.location.origin}/${uuid}`;
        showShareModal = true;
      })
      .catch((err) => {
        error = err.message;
        console.error('Save error:', err);
      })
      .finally(() => {
        saving = false;
      });
  }

  function closeShareModal() {
    showShareModal = false;
  }

  function resetToDefault() {
    kinkList = getDefaultTemplate();
    // Clear URL to root
    window.history.pushState({}, '', '/');
  }

  function updateNickname(value) {
    if (kinkList) {
      kinkList.nickname = value;
    }
  }

  function updateBlock(blockIndex, updatedBlock) {
    if (kinkList) {
      kinkList.blocks = kinkList.blocks.map((b, i) => i === blockIndex ? updatedBlock : b);
    }
  }

  function addBlock() {
    if (kinkList) {
      const newBlock = {
        id: `b_new_${Date.now()}`,
        title: 'New Block',
        comment: '',
        groups: [{ id: `g_new_${Date.now()}`, label: 'General' }],
        questions: []
      };
      kinkList.blocks = [...kinkList.blocks, newBlock];
    }
  }

  function deleteBlock(blockIndex) {
    if (kinkList) {
      kinkList.blocks = kinkList.blocks.filter((_, i) => i !== blockIndex);
    }
  }

  function openRatingsModal() {
    showRatingsModal = true;
  }

  function closeRatingsModal() {
    showRatingsModal = false;
  }

  function handleRatingsSave(event) {
    if (kinkList) {
      kinkList.ratings = event.detail;
    }
  }

  function openBlockModal(event) {
    const block = event.detail;
    const blockIndex = kinkList.blocks.findIndex(b => b.id === block.id);
    editingBlock = block;
    editingBlockIndex = blockIndex;
    showBlockModal = true;
  }

  function closeBlockModal() {
    showBlockModal = false;
    editingBlock = null;
    editingBlockIndex = -1;
  }

  function handleBlockSave(event) {
    if (kinkList && editingBlockIndex >= 0) {
      kinkList.blocks = kinkList.blocks.map((b, i) => i === editingBlockIndex ? event.detail : b);
    }
  }

  function closeWelcomeModal() {
    showWelcomeModal = false;
    if (typeof window !== 'undefined') {
      localStorage.setItem('kinklist_welcome_seen', 'true');
    }
  }
</script>

<div class="main">
  {#if loading}
    <div>Loading...</div>
  {:else if error && !kinkList}
    <div>Error: {error}</div>
  {:else if kinkList}
    <!-- Header -->
    <div class="header sticky-header">
      <div class="header-left">
        <div class="logo-svg" on:click={resetToDefault} role="button" tabindex="0" on:keydown={(e) => e.key === 'Enter' && resetToDefault()}>
          <img src="/img/ghost.svg" alt="Logo" width="80" height="68" />
        </div>
        <div class="kinklist-title" on:click={resetToDefault} role="button" tabindex="0" on:keydown={(e) => e.key === 'Enter' && resetToDefault()}>
          <span>Kinklist</span>
        </div>
        <div class="name-field">
          <input
            type="text"
            class="name-input"
            placeholder="Enter your name…"
            value={kinkList.nickname}
            on:input={(e) => updateNickname(e.target.value)}
          />
        </div>
      </div>
      <div class="header-right">
        <div class="rating-board">
          {#each kinkList.ratings as rating}
            <div class="rating-item">
              <div class="rating-color-indicator" style="background-color: {rating.color};"></div>
              <span class="rating-label">{rating.label}</span>
            </div>
          {/each}
          <button class="edit-ratings-btn" on:click={openRatingsModal} title="Edit ratings">
            <img src="/img/pencil.svg" alt="Edit"/>
          </button>
        </div>
        <button class="save-button" on:click={handleSave} disabled={saving}>
          {saving ? 'Saving...' : 'Save'}
        </button>
  </div>
  </div>

    <!-- Blocks container -->
    <div class="blocks-container">
      {#each kinkList.blocks as block, blockIndex}
        <Block
          {block}
          ratings={kinkList.ratings}
          on:update={(e) => updateBlock(blockIndex, e.detail)}
          on:delete={() => deleteBlock(blockIndex)}
          on:edit={openBlockModal}
        />
      {/each}
    </div>
    
    <!-- Fixed add block button -->
    <button class="add-block-btn" on:click={addBlock} title="Add new question block">
      <img src="/img/add-block.svg" alt="Add block" />
    </button>

    <!-- Edit Ratings Modal -->
    <EditRatingsModal
      ratings={kinkList.ratings}
      open={showRatingsModal}
      on:close={closeRatingsModal}
      on:save={handleRatingsSave}
    />

    <!-- Edit Block Modal -->
    <EditBlockModal
      block={editingBlock}
      open={showBlockModal}
      on:close={closeBlockModal}
      on:save={handleBlockSave}
    />

    <!-- Share Link Modal -->
    <ShareLinkModal
      open={showShareModal}
      link={shareLink}
      on:close={closeShareModal}
    />
  {/if}

  <!-- Welcome Modal -->
  <WelcomeModal
    open={showWelcomeModal}
    on:close={closeWelcomeModal}
  />
</div>

<style>
  .name-input {
    position: absolute;
    left: 27px;
    top: 17px;
    background: transparent;
    border: none;
    color: #7A80A2;
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    width: calc(100% - 54px);
    outline: none;
  }

  .name-input::placeholder {
    color: rgba(255, 255, 255, 0.5);
  }

  .save-button {
    padding: 10px 20px;
    background: #d7baff;
    border: none;
    border-radius: 25px;
    color: #121212;
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    cursor: pointer;
    transition: background 0.2s;
  }

  .save-button:hover:not(:disabled) {
    background: #c5a3ff;
  }

  .save-button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .sticky-header {
    top: 0;
    z-index: 100;
    background: #1b1c22;
    padding: 10px 0;
  }

  .add-block-btn {
    position: fixed;
    bottom: 30px;
    right: 30px;
    background: #373b4d;
    border: none;
    border-radius: 50%;
    cursor: pointer;
    padding: 12px;
    width: 56px;
    height: 56px;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
    transition: background 0.2s, transform 0.2s;
  }

  .add-block-btn:hover {
    background: #4a4f63;
    transform: scale(1.1);
  }

  .add-block-btn img {
    width: 24px;
    height: 24px;
    display: block;
  }

  .rating-board {
    display: flex;
    align-items: center;
    gap: 10px;
    background: #2c2e38;
    border-radius: 25px;
    padding: 8px 16px;
  }

  .rating-item {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .rating-color-indicator {
    width: 16px;
    height: 16px;
    border-radius: 4px;
    flex-shrink: 0;
  }

  .rating-label {
    font-family: 'Raleway', sans-serif;
    font-size: 14px;
    color: rgba(151, 156, 173, 1);
  }

  .edit-ratings-btn {
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: opacity 0.2s;
  }

  .edit-ratings-btn:hover {
    opacity: 0.8;
  }

  .edit-ratings-btn img {
    display: block;
  }

  @media (max-width: 768px) {
  .rating-board {
    flex-wrap: wrap;
    justify-content: flex-start;
    row-gap: 6px;
  }

  .rating-item {
    flex: 0 1 auto;
  }
}
</style>
