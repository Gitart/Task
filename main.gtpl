
{{define "lib"}}
   <h4>StarLight Brand Content</h4>
{{end}}


{{define "libs"}}
	  <meta charset="utf-8">

	  <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
	  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
	 
	  <!-- Latest compiled and minified JavaScript -->
	  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
	  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	  <script src="http://code.jquery.com/ui/1.10.4/jquery-ui.js"></script>
	  
	  	    <style>
						     .table-condensed>tbody>tr>td, 
						     .table-condensed>tbody>tr>th, 
						     .table-condensed>tfoot>tr>td, 
						     .table-condensed>tfoot>tr>th, 
						     .table-condensed>thead>tr>td, 
						     .table-condensed>thead>tr>th {padding: 2px; font-size:14px; margin:2px; vertical-align:middle; padding-left: 10px;}
			</style>        
{{end}}


{{define "header"}}
			<!-- Fixed navbar -->
			<div class="navbar navbar-inverse navbar-fixed-top" role="navigation">
				 <div class="container">
						<div class="navbar-header">
							<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
								   <span class="sr-only">Toggle navigation</span>
								   <span class="icon-bar"></span>
								   <span class="icon-bar"></span>
								   <span class="icon-bar"></span>
							</button>
								   <a class="navbar-brand" href="#">HEAD OFFICE</a>
						</div>
						
						<!--Menu-->
						<div class="navbar-collapse collapse">
						  
							<ul class="nav navbar-nav">
								  <li class="active"><a href="/api/system/pharmacy/view/done">Home</a></li>
								  <li><a href="/aboute">About</a></li>
								  <li><a href="#contact">Contact</a></li>
								
								<li class="dropdown">
								  <a href="#" class="dropdown-toggle" data-toggle="dropdown">Dropdown <b class="caret"></b></a>
								  <ul class="dropdown-menu">
									<li><a href="#">Action</a></li>
									<li><a href="#">Another action</a></li>
									<li><a href="#">Something else here</a></li>
									<li class="divider"></li>
									<li class="dropdown-header">Nav header</li>
									<li><a href="#">Separated link</a></li>
									<li><a href="#">One more separated link</a></li>
								  </ul>
								</li>
							  
							  <li class="dropdown">
								  <a href="#" class="dropdown-toggle" data-toggle="dropdown">Report<b class="caret"></b></a>
								  <ul class="dropdown-menu">
									<li><a href="#">Sales</a></li>
									<li><a href="#">Operation</a></li>
									<li><a href="#">View</a></li>
									<li class="divider"></li>
									<li class="dropdown-header">Log Files</li>
									<li><a href="#">More</a></li>
									<li><a href="#">My Link</a></li>
								  </ul>
								</li>
							</ul>
							
							<ul class="nav navbar-nav navbar-right">
								<li><a href="#"><span class="glyphicon glyphicon-user"></span> Sign Up</a></li>
								<li><a href="#"><span class="glyphicon glyphicon-log-in"></span> Login</a></li>
							</ul>
						</div>
						<!--/.nav-collapse -->
				</div>
		</div>
		<br>
		<br>
		<br>
{{end}}



{{define "footer"}}
		<footer>
		  <div>
			<p>&copy; 2014-2015 Softinform.
			<a href="http://morion.ua/licenses/by/3.0/" title="Creative Commons Attribution">Some rights reserved</a>;
				  please attribute properly and link back. Hosted by 
			<a href="http://morion.ua">&copy; Morion</a>.
			</p>
		  </div>
		</footer>
{{end}}

