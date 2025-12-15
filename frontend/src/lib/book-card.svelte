<script>
  export let Title;
  export let Description;
  export let BookCover;
  export let tags = [];
</script>

<div class="book-card">
  <div class="img">
    <img src={BookCover} alt="this is cover of {Title}" />
  </div>
  <div class="details">
    <div class="text">
      <h1 class="title">{Title}</h1>
      <p class="description">{Description}</p>
    </div>
    <div class="tags">
      {#each tags as tag}
        <div class="tag">
          {tag}
        </div>
      {/each}
    </div>
  </div>
</div>

<style>
  .book-card {
    width: 100svw;
    display: flex;
    margin-bottom: var(--space-md);

    .img {
      width: 10rem;
      height: 12rem;

      img {
        width: 10rem;
        height: 12rem;
        aspect-ratio: 3/4;
      }
    }
    .details {
      margin-inline: var(--space-sm);

      .text {
        height: 10rem;
        overflow: hidden;

        .title {
          font-size: var(--font-xlg);
          text-transform: capitalize;
        }
        .description {
          font-style: italic;
        }
      }
    }
    .tags {
      display: flex;
      .tag {
        background: var(--neutral-2);
        color: var(--text-inverse-color-1);
        padding-inline: var(--space-lg);
        padding-top: var(--space-sm);
        padding-bottom: var(--space-sm);
        border-radius: var(--space-md);
        margin-inline: var(--space-sm);
      }
      .tag:nth-child(1) {
        background: var(--blue-3);
        color: var(--text-color-1);
      }
    }
  }
  .book-card:hover {
    background: var(--bg2);
  }
</style>
