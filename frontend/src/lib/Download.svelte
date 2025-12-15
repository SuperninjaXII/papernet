<script>
  export let BookCover, Title, SubTitle, Author, Year, Description;
</script>

<section id="Download-card">
  <div class="top">
    <div class="img">
      <img src={BookCover} alt="cover img for {Title}" />
    </div>
    <div class="details">
      <h1 class="heading">{Title}</h1>
      <p class="sub-heading">{SubTitle}</p>
      <p class="author">{Author}</p>
      <p class="year">{Year}</p>
      <div class="tags"></div>
    </div>
  </div>
  <div class="bottom">
    <div class="description">
      <p>{Description}</p>
    </div>
    <button id="id">Download</button>
  </div>
</section>

<style>
  #Download-card {
    display: flex;
    flex-direction: column;
    width: 100svw;

    .top {
      display: flex;

      .img {
        width: 15rem;
        height: 20rem;
        overflow: hidden;

        img {
          width: 100%;
          height: 100%;
          aspect-ratio: 3/4;
        }
      }

      .details {
        margin-inline: var(--space-md);

        .heading {
          font-size: var(--font-xlg);
        }
        .sub-heading {
          font-size: var(--font-lg);
        }
        .author {
          color: var(--blue-1) !important;
          font-style: italic;
          font-weight: 200;
        }
        .year {
          color: var(--neutral-1) !important;
          font-style: oblique;
          font-weight: 600;
        }
      }
    }

    .bottom {
      display: flex;
      flex-direction: column;

      .description {
        max-height: 20rem;
        width: 70svw;
        margin-top: var(--space-md);
        margin-bottom: var(--space-md);
        margin-inline: auto;
      }
      button {
        width: 20svw;
        font-size: var(--font-md);
        background: var(--blue-3);
        color: var(--text-color-1);
        padding: var(--space-md);
        border: none;
        border-radius: var(--space-md);
        margin: auto;
      }
    }
  }
</style>
