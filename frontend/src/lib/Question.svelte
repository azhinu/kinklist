<script>
  import { createEventDispatcher } from 'svelte';
  import EditCommentModal from './EditCommentModal.svelte';

  export let question;
  export let ratings;
  export let groups;

  const dispatch = createEventDispatcher();
  
  let showCommentModal = false;
  let showCommentIcon = false;
  let showCommentTooltip = false;
  let tooltipPosition = { top: 0, right: 0 };
  let commentBtnElement;

  // Определяем, является ли устройство тач-устройством
  let isTouchDevice = false;
  
  if (typeof window !== 'undefined') {
    isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0;
    // На тач-устройствах иконка комментария всегда видима
    if (isTouchDevice) {
      showCommentIcon = true;
    }
  }

  function updateTooltipPosition() {
    if (commentBtnElement && typeof window !== 'undefined') {
      const rect = commentBtnElement.getBoundingClientRect();
      tooltipPosition = {
        top: rect.top,
        right: window.innerWidth - rect.right
      };
    }
  }

  function handleCommentBtnMouseEnter() {
    if (!isTouchDevice && question.comment && question.comment.trim()) {
      updateTooltipPosition();
      showCommentTooltip = true;
    }
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
  on:mouseenter={() => !isTouchDevice && (showCommentIcon = true)}
  on:mouseleave={() => !isTouchDevice && (showCommentIcon = false)}
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
    {@const hasComment = question.comment && question.comment.trim()}
    <div class="comment-btn-wrapper">
      <button 
        class="comment-btn" 
        bind:this={commentBtnElement}
        on:click={openCommentModal} 
        on:mouseenter={handleCommentBtnMouseEnter}
        on:mouseleave={() => !isTouchDevice && (showCommentTooltip = false)}
        title="View/Edit comment"
      >
        <img src={hasComment ? "/img/comment_active.svg" : "/img/comment.svg"} alt="Comment" width="24" height="24" />
      </button>
    </div>
  {/if}
</div>

<EditCommentModal
  comment={question.comment || ''}
  open={showCommentModal}
  on:close={closeCommentModal}
  on:save={(e) => handleCommentSave(e.detail)}
/>

{#if showCommentTooltip && question.comment && question.comment.trim()}
  <div 
    class="comment-tooltip" 
    style="top: {tooltipPosition.top}px; right: {tooltipPosition.right}px;"
    on:mouseenter={() => !isTouchDevice && (showCommentTooltip = true)}
    on:mouseleave={() => !isTouchDevice && (showCommentTooltip = false)}
  >
    <div class="comment-tooltip-content">
      {question.comment}
    </div>
  </div>
{/if}


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

  .comment-btn-wrapper {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1;
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
    z-index: 1;
  }

  .comment-btn:hover {
    opacity: 0.8;
  }

  .comment-btn:focus,
  .comment-btn:active {
    outline: none;
  }

  .comment-tooltip {
    position: fixed;
    z-index: 99999;
    pointer-events: auto;
    isolation: isolate;
    transform: translateY(calc(-100% - 8px));
  }

  .comment-tooltip-content {
    background: #2c2e38;
    border: 1px solid #7a80a2;
    border-radius: 8px;
    padding: 12px;
    max-width: 300px;
    min-width: 200px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
    color: rgba(151, 156, 173, 1);
    font-family: 'Raleway', sans-serif;
    font-size: 14px;
    line-height: 1.5;
    white-space: pre-wrap;
    word-wrap: break-word;
    position: relative;
    z-index: 99999;
  }

  .comment-tooltip-content::after {
    content: '';
    position: absolute;
    top: 100%;
    right: 16px;
    width: 0;
    height: 0;
    border-left: 8px solid transparent;
    border-right: 8px solid transparent;
    border-top: 8px solid #2c2e38;
  }

  .comment-tooltip-content::before {
    content: '';
    position: absolute;
    top: 100%;
    right: 15px;
    width: 0;
    height: 0;
    border-left: 9px solid transparent;
    border-right: 9px solid transparent;
    border-top: 9px solid #7a80a2;
    z-index: -1;
  }
</style>

