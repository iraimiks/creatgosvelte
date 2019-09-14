<div id="header">
  <div class="topnav">
        <div>
            <a class="top-nav-logo">Logo</a>
            <a class="active" href="/">Home</a>
            <a href="#news">News</a>
            <a href="#contact">Contact</a>
            <a href="/about">About</a>
        </div>
        <div class="topnav-right">
            <a>
            <i class="material-icons">
            shopping_cart
            </i>
            <span class="card-info">0</span>
            </a>
        </div>
  </div>
</div>
<script>

</script>
<style>
    #header{
        grid-area: header;
    }
    .topnav {
      overflow: hidden;
      background-color: #131313;
    }
    .topnav #myLinks {
      display: none;
    }
    .top-nav-logo{
        float: left;
        color: #f2f2f2;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 17px;
    }
    .topnav a {
      float: left;
      color: #f2f2f2;
      text-align: center;
      padding: 14px 16px;
      text-decoration: none;
      font-size: 17px;
    }
    .topnav a:hover {
      background-color: #480012;
      color: #b7ffb7;
    }
    .topnav-right {
     float: right;
    }
    .topnav-right a{
         float: right;
         padding-top: 10px;
        }
    .card-info {
        padding: 0;
    }
</style>