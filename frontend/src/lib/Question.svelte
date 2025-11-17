<script>
  import { createEventDispatcher } from 'svelte';
  import EditCommentModal from './EditCommentModal.svelte';

  export let question;
  export let ratings;
  export let groups;

  const dispatch = createEventDispatcher();
  
  let showCommentModal = false;
  let showCommentIcon = false;

  // Определяем, является ли устройство тач-устройством
  let isTouchDevice = false;
  
  if (typeof window !== 'undefined') {
    isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0;
  }

  function handleRatingClick(groupId, ratingId) {
    const updatedAnswers = [...question.answers];
    const existingAnswerIndex = updatedAnswers.findIndex(a => a.groupId === groupId);
    
    if (existingAnswerIndex >= 0) {
      if (updatedAnswers[existingAnswerIndex].ratingId === ratingId) {
        // Remove answer if clicking the same rating
        updatedAnswers.splice(existingAnswerIndex, 1);
      } else {
        // Update existing answer
        updatedAnswers[existingAnswerIndex] = { groupId, ratingId };
      }
    } else {
      // Add new answer
      updatedAnswers.push({ groupId, ratingId });
    }

    const updatedQuestion = {
      ...question,
      answers: updatedAnswers
    };
    dispatch('update', updatedQuestion);
  }

  function getRatingForGroup(groupId) {
    const answer = question.answers.find(a => a.groupId === groupId);
    return answer ? answer.ratingId : null;
  }

  function handleTitleChange(e) {
    const updatedQuestion = {
      ...question,
      title: e.target.value
    };
    dispatch('update', updatedQuestion);
  }

  function openCommentModal() {
    showCommentModal = true;
  }

  function closeCommentModal() {
    showCommentModal = false;
  }

  function handleCommentSave(comment) {
    const updatedQuestion = {
      ...question,
      comment: comment
    };
    dispatch('update', updatedQuestion);
  }

  function handleTitleClick(e) {
    // Если клик был по кнопке комментария, не обрабатываем
    if (e.target.closest('.comment-btn')) {
      return;
    }
  }

</script>

<div class="question-title"
  role="button"
  tabindex="0"
  on:mouseenter={() => (showCommentIcon = true)}
  on:mouseleave={() => (showCommentIcon = false)}
  on:click={handleTitleClick}
  on:keydown={(e) => e.key === 'Enter' && handleTitleClick(e)}
>
  <div class="text-node-html">
    <div class="root rich-text root-0">
      <div class="paragraph-set root-0-paragraph-set-0">
        <p class="paragraph root-0-paragraph-set-0-paragraph-0" dir="ltr">
          {question.title}
        </p>
      </div>
    </div>
  </div>
  {#if showCommentIcon || (question.comment && question.comment.trim())}
    <button class="comment-btn" on:click={openCommentModal} title="View/Edit comment">
      <img src="/img/comment.svg" alt="Comment" width="24" height="24" />
    </button>
  {/if}
</div>

<EditCommentModal
  comment={question.comment || ''}
  open={showCommentModal}
  on:close={closeCommentModal}
  on:save={(e) => handleCommentSave(e.detail)}
/>


{#each groups as group}
  <div class="rates-container">
    <div class="group-box">
      <span class="group-label">{group.label}</span>
    </div>
    <div class="rates">
      {#each ratings as rating}
        {@const isSelected = getRatingForGroup(group.id) === rating.id}
        <button
          class="rate-color"
          class:selected={isSelected}
          style="background-color: {rating.color};"
          on:click={() => handleRatingClick(group.id, rating.id)}
          title={rating.label}
        ></button>
      {/each}
    </div>
  </div>
{/each}

<style>
  .question-title {
    position: relative;
    width: 100%;
    margin-bottom: 10px;
    display: flex;
    align-items: center;
    gap: 10px;
    justify-content: flex-end;
  }

  .text-node-html {
    flex: 1;
    display: flex;
    justify-content: flex-end;
  }

  .root {
    display: flex;
    justify-content: flex-end;
  }

  .paragraph-set {
    display: inline-flex;
    flex-direction: column;
    min-width: 100%;
    align-items: flex-end;
  }

  .paragraph {
    margin: 0;
    text-align: right;
    width: 100%;
  }

  .group-box {
    position: relative;
    background: #7a80a2;
    border-radius: 15px;
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 10px;
    margin-right: 10px;
  }

  .group-label {
    font-family: 'Raleway', sans-serif;
    font-size: 14px;
    color: white;
  }

  .rates {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 5px;
    flex-direction: row-reverse;
  }

  .rate-color {
    width: 20px;
    height: 20px;
    border-radius: 50%;
    border: 2px solid transparent;
    cursor: pointer;
    transition: border-color 0.2s;
    padding: 0;
    margin: 0;
  }

  .rate-color.selected {
    border-color: #d7baff;
    border-width: 2px;
  }

  .rate-color:hover {
    border-color: #d7baff;
    border-width: 2px;
  }

  .comment-btn {
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

  .comment-btn:hover {
    opacity: 0.8;
  }

  .comment-btn:focus,
  .comment-btn:active {
    outline: none;
  }
</style>

