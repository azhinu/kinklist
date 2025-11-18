<script>
  import { createEventDispatcher } from 'svelte';
  import { onDestroy, tick } from 'svelte';

  export let comment = '';
  export let open = false;

  const dispatch = createEventDispatcher();
  
  let editedComment = '';
  let textareaElement;

  function handleEscape(e) {
    if (e.key === 'Escape' && open) {
      e.preventDefault();
      e.stopPropagation();
      handleClose();
    }
  }

  $: if (typeof window !== 'undefined' && open) {
    window.addEventListener('keydown', handleEscape, true);
  } else if (typeof window !== 'undefined') {
    window.removeEventListener('keydown', handleEscape, true);
  }

  onDestroy(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('keydown', handleEscape, true);
    }
  });

  $: if (open) {
    editedComment = comment || '';
    // Фокусируем текстовое поле при открытии модального окна
    if (typeof window !== 'undefined') {
      // Используем tick() для ожидания обновления DOM
      tick().then(() => {
        textareaElement?.focus();
      });
    }
  }

  function handleClose() {
    dispatch('close');
  }

  function handleSave() {
    dispatch('save', editedComment);
    handleClose();
  }

  function updateComment(value) {
    editedComment = value;
  }
</script>

{#if open}
  <div class="modal-overlay" role="dialog" aria-modal="true" tabindex="-1" on:click={handleClose} on:keydown={(e) => e.key === 'Escape' && handleClose()}>
    <div class="modal-content" role="document" on:click|stopPropagation on:keydown={(e) => {
      if (e.key !== 'Escape') {
        e.stopPropagation();
      } else {
        handleClose();
      }
    }}>
      <div class="update-comment-block">
        <!-- Comment edit -->
        <div class="comment-edit">
          <!-- Comment edit title -->
          <div class="comment-edit-title">
            <span class="comment-edit-title-text">View/Edit comment:</span>
          </div>

          <!-- Comment input -->
          <div class="comment-input-container">
            <textarea
              bind:this={textareaElement}
              class="comment-text"
              value={editedComment}
              on:input={(e) => updateComment(e.target.value)}
              on:keydown={(e) => e.key === 'Escape' && handleClose()}
              placeholder="Enter comment"
              rows="4"
            ></textarea>
          </div>
        </div>

        <!-- Tools -->
        <div class="tools">
          <button class="save-btn" on:click={handleSave} title="Save">
            <img src="/img/save_btn.svg" alt="Save"/>
          </button>
          <button class="close-btn" on:click={handleClose} title="Close">
            <img src="/img/close_btn.svg" alt="Close"/>
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

  .update-comment-block {
    position: relative;
    width: 100%;
    background: #1b1c22;
    border-radius: 25px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 20px;
    padding-bottom: 80px;
    min-width: 400px;
    max-width: 600px;
  }

  @media (max-width: 420px) {
    .update-comment-block {
      min-width: 280px;
    }
  }

  .comment-edit {
    position: relative;
    width: auto;
    height: auto;
    overflow: hidden;
    display: grid;
    align-items: start;
    gap: 10px;
    grid-template-columns: 1fr;
    min-height: 100px;
    min-width: 251px;
    margin-bottom: 0;
  }

  .comment-edit-title {
    position: relative;
    width: auto;
    height: auto;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 5px;
    padding: 5px;
    min-height: 34px;
    min-width: 125px;
  }

  .comment-edit-title-text {
    background: transparent;
    border: none;
    color: inherit;
    font-family: 'Raleway', sans-serif;
    font-size: 18px;
    font-weight: bold;
    outline: none;
    min-width: 0;
    overflow-wrap: anywhere;
  }

  .comment-input-container {
    position: relative;
    width: auto;
    height: auto;
    background: #373b4d;
    border-radius: 25px;
    display: flex;
    align-items: flex-start;
    padding: 17px 27px;
    min-width: 150px;
  }

  .comment-text {
    background: transparent;
    border: none;
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    width: 100%;
    outline: none;
    resize: vertical;
    min-height: 80px;
  }

  .comment-text::placeholder {
    color: rgba(151, 156, 173, 0.5);
  }

  .tools {
    position: absolute;
    right: 20px;
    bottom: 20px;
    width: auto;
    height: auto;
    background: #373b4d;
    border-radius: 25px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    padding: 5px;
    z-index: 10;
  }

  .save-btn,
  .close-btn {
    position: relative;
    width: 32px;
    height: 32px;
    flex-shrink: 0;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .save-btn:hover,
  .close-btn:hover {
    opacity: 0.8;
  }

  .save-btn:focus,
  .close-btn:focus,
  .save-btn:active,
  .close-btn:active {
    outline: none;
  }
</style>

