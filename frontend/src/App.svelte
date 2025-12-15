<script>
  import Book from "./lib/Book.svelte";
  import Header from "./lib/Header.svelte";
  import BookCard from "./lib/book-card.svelte";
  import Download from "./lib/Download.svelte";
  let tags = ["test", "tag2"];
  fetch("/api/books")
    .then((Response) => Response.text())
    .then((data) => console.log(data));
</script>

<Header></Header>
<main>
  <Download
    BookCover="/test.png"
    Title="go programming"
    SubTitle="for dummies"
    Description="fgfdjh dfdsfdf fdfdf"
    Author="wei meng li"
    Year="2025"
  ></Download>
  <section id="search-list">
    <BookCard
      Title="test"
      Description="lorem ipsum gjlgkjlf lglfkngl dlfgjdlkfjdlkfjl ,fnlsfnlsnflksndf fhflk ldkflhfhdkjlfh lfhkslkfhsf kjfsfbjsbf sfklsfjhsf kjsfkhsfkjslfnlfhklshnslkfhslkfjh fbnfnl fbdsfbdljfbkljf lnfldkflsdf"
      BookCover="/test.png"
      {tags}
    ></BookCard>
    <BookCard
      Title="test"
      Description="lorem ipsum gjlgkjlf lglfkngl dlfgjdlkfjdlkfjl ,fnlsfnlsnflksndf fhflk ldkflhfhdkjlfh lfhkslkfhsf kjfsfbjsbf sfklsfjhsf kjsfkhsfkjslfnlfhklshnslkfhslkfjh fbnfnl fbdsfbdljfbkljf lnfldkflsdf"
      BookCover="/test.png"
      {tags}
    ></BookCard>
  </section>
  <h2 class="heading">Latest Books</h2>
  <section id="latest-books">
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
    <Book Title="this is the book on bla bla" BookCover="/test.png"></Book>
  </section>
</main>

<style>
  #latest-books {
    display: flex;
    flex-wrap: wrap;
    max-width: 80svw;
    margin: auto;
  }
  .heading {
    text-align: center;
    margin-top: var(--space-xlg);
  }
</style>
