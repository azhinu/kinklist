<script>
  import { createEventDispatcher } from 'svelte';

  export let block = null;
  export let open = false;

  const dispatch = createEventDispatcher();
  
  let editedBlock = null;

  $: if (open && block) {
    editedBlock = {
      ...block,
      groups: block.groups ? block.groups.map(g => ({ ...g })) : [],
      questions: block.questions ? block.questions.map(q => ({ ...q })) : []
    };
  }

  function handleClose() {
    dispatch('close');
  }

  function handleSave() {
    dispatch('save', editedBlock);
    handleClose();
  }

  function updateBlockTitle(value) {
    if (editedBlock) {
      editedBlock.title = value;
      editedBlock = editedBlock; // trigger reactivity
    }
  }

  function handleBlockNameFocus(e) {
    if (e.target.value === 'New Block') {
      e.target.value = '';
      updateBlockTitle('');
    }
  }

  function addGroup() {
    if (editedBlock) {
      const newGroup = {
        id: `g_new_${Date.now()}`,
        label: ''
      };
      editedBlock.groups = [...editedBlock.groups, newGroup];
    }
  }

  function removeGroup(index) {
    if (editedBlock) {
      editedBlock.groups = editedBlock.groups.filter((_, i) => i !== index);
    }
  }

  function updateGroupLabel(index, value) {
    if (editedBlock) {
      editedBlock.groups[index].label = value;
      editedBlock = editedBlock; // trigger reactivity
    }
  }

  function addQuestion() {
    if (editedBlock) {
      const newQuestion = {
        id: `q_new_${Date.now()}`,
        title: '',
        comment: '',
        answers: []
      };
      editedBlock.questions = [...editedBlock.questions, newQuestion];
    }
  }

  function removeQuestion(index) {
    if (editedBlock) {
      editedBlock.questions = editedBlock.questions.filter((_, i) => i !== index);
    }
  }

  function updateQuestionTitle(index, value) {
    if (editedBlock) {
      editedBlock.questions[index].title = value;
      editedBlock = editedBlock; // trigger reactivity
    }
  }
</script>

{#if open && editedBlock}
  <div class="modal-overlay" role="dialog" aria-modal="true" on:click={handleClose} on:keydown={(e) => e.key === 'Escape' && handleClose()}>
    <div class="modal-content" role="document" on:click|stopPropagation>
      <div class="update-question-block">
        <!-- Block name -->
        <div class="block-name">
          <input
            type="text"
            class="block-name-input"
            value={editedBlock.title}
            on:input={(e) => updateBlockTitle(e.target.value)}
            on:focus={handleBlockNameFocus}
            placeholder="Enter block name"
          />
        </div>

        <div class="block-edit-content">
          <!-- Groups edit -->
          <div class="groups-edit">
            <!-- Groups edit title -->
            <div class="groups-edit-title">
              <span class="groups-edit-title-text">Groups:</span>
              <button class="add-icon" on:click={addGroup} title="Add group">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <circle cx="12" cy="12" r="10" stroke="#7885cd" stroke-width="2"/>
                  <path d="M12 8V16M8 12H16" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
                </svg>
              </button>
            </div>

            <!-- Group blocks -->
            {#each editedBlock.groups as group, index}
              <div class="group-block">
                <div class="group-title-input">
                  <input
                    type="text"
                    class="group-title-text"
                    value={group.label}
                    on:input={(e) => updateGroupLabel(index, e.target.value)}
                    placeholder="Enter block name"
                  />
                </div>
                <button class="remove-icon" on:click={() => removeGroup(index)} title="Remove group">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <circle cx="12" cy="12" r="10" stroke="#7885cd" stroke-width="2"/>
                    <path d="M8 12H16" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
                  </svg>
                </button>
              </div>
            {/each}
          </div>

          <!-- Questions edit -->
          <div class="questions-edit">
            <!-- Questions edit title -->
            <div class="questions-edit-title">
              <span class="questions-edit-title-text">Questions:</span>
              <button class="add-icon" on:click={addQuestion} title="Add question">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <circle cx="12" cy="12" r="10" stroke="#7885cd" stroke-width="2"/>
                  <path d="M12 8V16M8 12H16" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
                </svg>
              </button>
            </div>

            <!-- Question blocks -->
            {#each editedBlock.questions as question, index}
              <div class="question-block">
                <div class="question-title-input">
                  <input
                    type="text"
                    class="question-title-text"
                    value={question.title}
                    on:input={(e) => updateQuestionTitle(index, e.target.value)}
                    placeholder="Enter the question"
                  />
                </div>
                <button class="remove-icon" on:click={() => removeQuestion(index)} title="Remove question">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <circle cx="12" cy="12" r="10" stroke="#7885cd" stroke-width="2"/>
                    <path d="M8 12H16" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
                  </svg>
                </button>
              </div>
            {/each}
          </div>
        </div>

        <!-- Tools -->
        <div class="tools">
          <button class="save-btn" on:click={handleSave} title="Save">
            <svg width="32" height="32" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="16" cy="16" r="13.33" stroke="#7885cd" stroke-width="2"/>
              <path d="M11 16L14.67 19.67L21 13.33" stroke="#7885cd" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
          <button class="close-btn" on:click={handleClose} title="Close">
            <svg width="32" height="32" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M8 8L24 24M8 24L24 8" stroke="#7885cd" stroke-width="2" stroke-linecap="round"/>
            </svg>
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

  .update-question-block {
    position: relative;
    width: 100%;
    background: #1b1c22;
    border-radius: 25px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 20px;
    padding-bottom: 80px;
    min-width: 500px;
    max-width: 800px;
  }

    @media (max-width: 500px) {
    .update-question-block {
      min-width: auto;
      max-width: 100%;
    }
  }


  .block-name {
    position: relative;
    width: auto;
    height: 50px;
    background: #373b4d;
    border-radius: 25px;
    display: flex;
    align-items: center;
    padding: 17px 27px;
  }

  .block-name-input {
    background: transparent;
    border: none;
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    width: 100%;
    outline: none;
  }

  .block-name-input::placeholder {
    color: rgba(151, 156, 173, 0.5);
  }

  .block-edit-content {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
    align-items: start;
  }

    @media (max-width: 500px) {
    .block-edit-content {
      grid-template-columns: 1fr;
    }
  }


  .groups-edit,
  .questions-edit {
    position: relative;
    width: auto;
    height: auto;
    overflow: hidden;
    display: grid;
    align-items: start;
    gap: 10px;
    grid-template-columns: 1fr;
    min-height: 100px;
  }

  .groups-edit-title,
  .questions-edit-title {
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
  }

  .groups-edit-title-text,
  .questions-edit-title-text {
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

  .add-icon:hover {
    opacity: 0.8;
  }

  .group-block,
  .question-block {
    position: relative;
    width: auto;
    height: auto;
    display: flex;
    align-items: center;
    gap: 5px;
  }

  .group-title-input,
  .question-title-input {
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

  .group-title-text,
  .question-title-text {
    background: transparent;
    border: none;
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 16px;
    width: 100%;
    outline: none;
  }

  .group-title-text::placeholder,
  .question-title-text::placeholder {
    color: rgba(151, 156, 173, 0.5);
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

