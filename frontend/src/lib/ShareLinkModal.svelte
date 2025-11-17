<script>
  import { createEventDispatcher } from 'svelte';
  import { onDestroy } from 'svelte';

  export let open = false;
  export let link = '';

  const dispatch = createEventDispatcher();
  let isClicked = false;

  function handleEscape(e) {
    if (e.key === 'Escape' && open) {
      handleClose();
    }
  }

  $: if (typeof window !== 'undefined' && open) {
    window.addEventListener('keydown', handleEscape);
  } else if (typeof window !== 'undefined') {
    window.removeEventListener('keydown', handleEscape);
  }

  onDestroy(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('keydown', handleEscape);
    }
  });

  function handleClose() {
    dispatch('close');
  }

  function copyToClipboard() {
    if (link) {
      isClicked = true;
      navigator.clipboard.writeText(link).then(() => {
        // Можно добавить уведомление об успешном копировании
        setTimeout(() => {
          isClicked = false;
        }, 200);
      }).catch(err => {
        console.error('Failed to copy:', err);
        isClicked = false;
      });
    }
  }
</script>

{#if open}
  <div class="modal-overlay" role="dialog" aria-modal="true" on:click={handleClose} on:keydown={(e) => e.key === 'Escape' && handleClose()}>
    <div class="modal-content" role="document" on:click|stopPropagation>
      <div class="get-link">
        <span class="share-text">You can share results with this link:</span>
        
        <div class="gen-link-block">
          <div class="link-container">
            <span class="link-text" class:clicked={isClicked} on:click={copyToClipboard} role="button" tabindex="0" on:keydown={(e) => e.key === 'Enter' && copyToClipboard()}>{link || 'Link'}</span>
            <button class="copy-btn" on:click={copyToClipboard} title="Copy link">
              <img src="/img/copy_btn.svg" alt="Copy" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2000;
  }

  .modal-content {
    position: relative;
    max-width: 90vw;
    max-height: 90vh;
    overflow: auto;
  }

  .get-link {
    position: relative;
    width: 100%;
    background: #1b1c22;
    border-radius: 25px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 20px;
    padding: 20px;
    flex-direction: column;
    max-width: 600px;
  }

  .share-text {
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    font-weight: 400;
    height: 20px;
    text-align: center;
  }

  .gen-link-block {
    position: relative;
    width: auto;
    height: auto;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 5px;
    width: 100%;
  }

  .link-container {
    position: relative;
    width: auto;
    height: 50px;
    background: #373b4d;
    border-radius: 25px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 16px;
    padding: 16px;
    flex: 1;
    min-width: 0;
  }

  .link-text {
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    font-weight: 400;
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    min-width: 0;
    cursor: pointer;
    user-select: none;
    transition: color 0.2s;
  }

  .link-text:hover {
    opacity: 0.8;
  }

  .link-text.clicked {
    color: #687DF7;
  }

  .copy-btn {
    position: relative;
    width: 25px;
    height: 25px;
    flex-shrink: 0;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    outline: none;
  }

  .copy-btn:focus {
    outline: none;
  }

  .copy-btn:active {
    outline: none;
  }

  .copy-btn:hover {
    opacity: 0.8;
  }

  .copy-btn img {
    display: block;
    width: 25px;
    height: 25px;
  }
</style>

