<script>
  export let Title;
  export let BookCover;
</script>

<div class="book">
  <img src={BookCover} alt="cover image of {Title}" />
  <p class="title">{Title}</p>
</div>

<style>
  .book {
    width: 10rem;
    height: 15rem;
    background: var(--bg);
    overflow: hidden;
    margin: 2rem;

    img {
      width: 10rem;
      border-radius: 0.5rem;
    }
    .title {
      font-size: 1.2rem;
      color: #fff;
      text-align: center;
    }
  }
  /*mobile*/
  @media (width <=600px) {
    .book {
      width: 5rem;
      height: 8rem;
      img {
        width: 5rem;
      }
    }
  }
</style>
