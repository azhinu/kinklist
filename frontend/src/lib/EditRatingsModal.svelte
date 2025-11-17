<script>
  import { createEventDispatcher } from 'svelte';
  import { onDestroy } from 'svelte';

  export let ratings = [];
  export let open = false;

  const dispatch = createEventDispatcher();
  
  let editedRatings = [];

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

  // Normalize color format: remove alpha channel if present (e.g., #5354b6FF -> #5354b6)
  function normalizeColor(color) {
    if (!color) return '#979cad';
    // Ensure it starts with #
    if (!color.startsWith('#')) {
      color = '#' + color;
    }
    // If color has 9 characters (with alpha), remove last 2
    if (color.length === 9) {
      return color.substring(0, 7);
    }
    // If color has 7 characters (without alpha), return as is
    if (color.length === 7) {
      return color;
    }
    // Fallback
    return '#979cad';
  }

  // Ensure color has alpha channel for storage (e.g., #5354b6 -> #5354b6FF)
  function ensureAlphaColor(color) {
    if (!color) return '#979cadFF';
    // If color has 7 characters (without alpha), add FF
    if (color.length === 7 && color.startsWith('#')) {
      return color + 'FF';
    }
    // If already has alpha, return as is
    if (color.length === 9 && color.startsWith('#')) {
      return color;
    }
    return color;
  }

  $: if (open && ratings) {
    editedRatings = ratings.map(r => ({ 
      ...r, 
      color: normalizeColor(r.color) // Normalize for display
    }));
  }

  function handleClose() {
    dispatch('close');
  }

  function handleSave() {
    // Ensure colors have alpha channel before saving
    const ratingsToSave = editedRatings.map(r => ({
      ...r,
      color: ensureAlphaColor(r.color)
    }));
    dispatch('save', ratingsToSave);
    handleClose();
  }

  function addRating() {
    const newRating = {
      id: `r_new_${Date.now()}`,
      label: '',
      color: '#979cad'
    };
    editedRatings = [...editedRatings, newRating];
  }

  function removeRating(index) {
    editedRatings = editedRatings.filter((_, i) => i !== index);
  }

  function updateRatingLabel(index, value) {
    editedRatings[index].label = value;
    editedRatings = editedRatings; // trigger reactivity
  }

  function updateRatingColor(index, color) {
    // Color from input is already in #RRGGBB format (no alpha)
    editedRatings[index].color = normalizeColor(color);
    editedRatings = editedRatings; // trigger reactivity
  }
</script>

{#if open}
  <div class="modal-overlay" role="dialog" aria-modal="true" on:click={handleClose} on:keydown={(e) => e.key === 'Escape' && handleClose()}>
    <div class="modal-content" role="document" on:click|stopPropagation>
      <div class="update-ratings-block">
        <!-- Ratings edit -->
        <div class="ratings-edit">
          <!-- Ratings edit title -->
          <div class="ratings-edit-title">
            <span class="ratings-edit-title-text">Edit ratings:</span>
          </div>

          <!-- Rating blocks -->
          {#each editedRatings as rating, index}
            <div class="rating-block">
              <div class="rating-title-input">
                <input
                  type="text"
                  class="rating-title-text"
                  value={rating.label}
                  on:input={(e) => updateRatingLabel(index, e.target.value)}
                  placeholder="Enter block name"
                />
              </div>
              <input
                type="color"
                class="rating-color"
                value={rating.color}
                on:input={(e) => updateRatingColor(index, e.target.value)}
              />
              <button class="remove-icon" on:click={() => removeRating(index)} title="Remove rating">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <circle cx="12" cy="12" r="10" stroke="#7885cd" stroke-width="2"/>
                  <path d="M8 12H16" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
                </svg>
              </button>
            </div>
          {/each}
        </div>

        <!-- Tools -->
        <div class="tools">
          <button class="add-icon" on:click={addRating} title="Add rating">
            <img src="/img/add_btn.svg" alt="Add"/>
          </button>
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

  .update-ratings-block {
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
    .update-ratings-block {
      min-width: 280px;
    }
  }

  .ratings-edit {
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

  .ratings-edit-title {
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

  .ratings-edit-title-text {
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

  .add-icon {
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

  .add-icon:hover {
    opacity: 0.8;
  }

  .rating-block {
    position: relative;
    width: auto;
    height: auto;
    display: flex;
    align-items: center;
    gap: 5px;
  }

  .rating-title-input {
    position: relative;
    width: auto;
    height: 50px;
    background: #373b4d;
    border-radius: 25px;
    display: flex;
    align-items: center;
    padding: 17px 27px;
    flex: 1;
    min-width: 150px;
  }

  .rating-title-text {
    background: transparent;
    border: none;
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    width: 100%;
    outline: none;
  }

  .rating-title-text::placeholder {
    color: rgba(151, 156, 173, 0.5);
  }

  .rating-color {
    width: 32px;
    height: 32px;
    border: none;
    border-radius: 50%;
    cursor: pointer;
    flex-shrink: 0;
    padding: 0;
    background: transparent;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    overflow: hidden;
  }

  .rating-color::-webkit-color-swatch-wrapper {
    padding: 0;
  }

  .rating-color::-webkit-color-swatch {
    border: none;
    border-radius: 50%;
  }

  .rating-color::-moz-color-swatch {
    border: none;
    border-radius: 50%;
  }

  .remove-icon {
    position: relative;
    width: 24px;
    height: 24px;
    flex-shrink: 0;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .remove-icon:hover {
    opacity: 0.8;
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
</style>

