<script>
  import { createEventDispatcher } from 'svelte';
  import { onDestroy } from 'svelte';

  export let open = false;

  const dispatch = createEventDispatcher();

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
</script>

{#if open}
  <div class="modal-overlay" role="dialog" aria-modal="true">
    <div class="modal-content" role="document" on:click|stopPropagation>
      <div class="welcome-block">
        <!-- Welcome section -->
        <div class="welcome-section">
          <h2 class="section-title">Welcome!</h2>
          <p class="section-text">
            You can fill out the default questionnaire or build your own.
            <br><br>
            Everything is optional and fully editable, including rating options, blocks, questions and comments.
            <br><br>
            Use <img src="/img/pencil.svg" alt="Edit" class="inline-icon" /> button to edit content.
            <br>
            Hover over a block or question title to reveal editing tools.
            <br>
            Click on selected rating to deselect it.
          </p>
        </div>

        <!-- Saving section -->
        <div class="saving-section">
          <h2 class="section-title">Saving</h2>
          <p class="section-text">
            <button class="save-button-example" disabled>
              Save
            </button> creates a new snapshot of what you see.
            <br>
            Each Save gives you a new link.
            <br>
            Older versions never change.
          </p>
        </div>

        <!-- PRs section -->
        <div class="prs-section">
          <p class="prs-text">
            <a href="https://github.com/azhinu/kinklist" target="_blank" rel="noopener noreferrer" class="github-link">
              <img src="/img/github_icon.svg" alt="GitHub" class="github-icon" />
              PRs are welcome
            </a>
          </p>
        </div>

        <!-- Got it button -->
        <div class="got-it-container">
          <button class="got-it-button" on:click={handleClose}>
            Got it
          </button>
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

  .welcome-block {
    position: relative;
    width: 100%;
    background: #1b1c22;
    border-radius: 25px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 30px;
    min-width: 400px;
    max-width: 600px;
  }

  @media (max-width: 420px) {
    .welcome-block {
      min-width: 280px;
      padding: 20px;
    }
  }

  .welcome-section,
  .saving-section {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .section-title {
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 18px;
    font-weight: bold;
    margin: 0;
  }

  .section-text {
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    line-height: 1.5;
    margin: 0;
  }

  .inline-icon {
    width: 16px;
    height: 16px;
    vertical-align: middle;
    margin: 0 4px;
    display: inline-block;
  }

  .save-button-example {
    background: #d7baff;
    border: none;
    border-radius: 25px;
    color: #121212;
    font-family: 'Raleway', sans-serif;
    font-size: 10px;
    cursor: default;
    display: inline-block;
    vertical-align: middle;
  }

  .prs-section {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 10px 0;
  }

  .prs-text {
    margin: 0;
  }

  .github-link {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    text-decoration: none;
    transition: opacity 0.2s;
  }

  .github-link:hover {
    opacity: 0.8;
  }

  .github-icon {
    width: 20px;
    height: 20px;
    display: block;
  }

  .got-it-container {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 10px;
  }

  .got-it-button {
    padding: 10px 30px;
    background: #d7baff;
    border: none;
    border-radius: 25px;
    color: #121212;
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    cursor: pointer;
    transition: background 0.2s;
  }

  .got-it-button:hover {
    background: #c5a3ff;
  }

  .got-it-button:focus {
    outline: none;
  }
</style>

