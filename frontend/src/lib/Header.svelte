<script>
</script>

<header>
  <div class="top-row">
    <div class="brand">papernet</div>
  </div>
  <div class="search-container">
    <input type="text" id="search-bar" placeholder="what do you need ?" />
    <button>search</button>
  </div>
</header>

<style>
  @import url("../style/variables.css");
  header {
    display: flex;
    flex-direction: column;
    width: 100svw;
    position: sticky;
    z-index: 2;
    top: 0;
    background: transparent;

    .top-row {
      display: flex;
      background: transparent;
      backdrop-filter: blur(10px);

      .brand {
        background: transparent;
        backdrop-filter: blur(50px);
        margin: auto;
        font-size: var(--font-xlg);
        color: var(--text-color-1);
      }
    }
    .search-container {
      background: transparent;
      display: flex;
      margin-inline: auto;
      margin-top: var(--space-sm);
      input {
        background: var(--bg2-transparent);
        backdrop-filter: blur(10px);
        border: none;
        width: 20svw;
        height: 5svh;
        padding: var(--space-sm);
        border-radius: var(--space-md);
        color: var(--text-color-1);
        text-align: center;
        font-size: var(--font-md);
      }
      button {
        padding-inline: var(--space-lg);
        border: none;
        border-radius: var(--space-md);
        color: var(--text-color-1);
        background: var(--blue-3);
        margin-inline: var(--space-sm);
        font-size: var(--font-md);
      }
    }
  }
  /*mobile*/
  @media (width <=600px) {
    input {
      width: 50svw !important;
      height: 2rem !important;
    }
  }
</style>
