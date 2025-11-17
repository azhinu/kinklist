<script>
  import Question from './Question.svelte';
  import { createEventDispatcher } from 'svelte';

  export let block;
  export let ratings;

  const dispatch = createEventDispatcher();
  let showEditIcon = false;
  
  // Определяем, является ли устройство тач-устройством
  let isTouchDevice = false;
  
  if (typeof window !== 'undefined') {
    isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0;
  }

  function handleEditClick() {
    dispatch('edit', block);
  }

  function handleDeleteClick() {
    dispatch('delete');
  }
  
  function handleTitleClick(e) {
    // На тач-устройствах переключаем видимость кнопок
    if (isTouchDevice) {
      // Если клик был по области с кнопками, не переключаем
      if (e.target.closest('.block-title-icons')) {
        return;
      }
      showEditIcon = !showEditIcon;
    } else {
      // На устройствах без мышки, если кнопки не видны, игнорируем тап
      if (!showEditIcon) {
        e.preventDefault();
        e.stopPropagation();
        return;
      }
    }
  }

  function handleUpdate(questionIndex, updatedQuestion) {
    const updatedBlock = {
      ...block,
      questions: block.questions.map((q, i) => i === questionIndex ? updatedQuestion : q)
    };
    dispatch('update', updatedBlock);
  }


  function handleTitleChange(e) {
    const updatedBlock = {
      ...block,
      title: e.target.value
    };
    dispatch('update', updatedBlock);
  }

</script>

<div class="block" role="region">
  <div class="block-title"
    role="button"
    tabindex="0"
    on:mouseenter={() => !isTouchDevice && (showEditIcon = true)}
    on:mouseleave={() => !isTouchDevice && (showEditIcon = false)}
    on:click={handleTitleClick}
    on:keydown={(e) => e.key === 'Enter' && handleTitleClick(e)}
    >
    <span class="block-title-text">{block.title}</span>
    {#if showEditIcon}
    <div class="block-title-icons">
      <button class="edit-icon" title="Edit block" on:click={handleEditClick}>
        <img src="/img/pencil.svg" alt="Edit"/>
      </button>
      <button class="remove-icon" title="Remove block" on:click={handleDeleteClick}>
        <img src="/img/remove.svg" alt="Remove"/>
      </button>
    </div>
    {/if}
  </div>

  <div class="question-container">
    {#each block.questions as question, questionIndex}
      <Question
        {question}
        {ratings}
        groups={block.groups}
        on:update={(e) => handleUpdate(questionIndex, e.detail)}
      />
    {/each}
  </div>
</div>

<style>
.block-title-text {
  background: transparent;
  border: none;
  color: inherit;
  font-family: 'Raleway', sans-serif;
  font-size: 18px;
  font-weight: bold;
  outline: none;
  grid-column: 1;
  grid-row: 1;
  min-width: 0;
  overflow-wrap: anywhere;
}

.edit-icon,
.remove-icon {
  background: transparent;
  border: none;
  padding: 0;
  margin: 0;
}
</style>

