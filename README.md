<h1 align="center">KCardCLI</h1>

<p align="center">
  KCardCLI is an interactive command-line tool to generate custom keycard commands for <strong>SCP: Secret Laboratory</strong> servers.<br/>
  It helps server admins and modders create valid and ready-to-use keycard command strings by guiding them through all necessary parameters.
</p>

<p align="center">
  <img src="https://repobeats.axiom.co/api/embed/7ac5eab2201227aa7bfd9d3d931b22fd0959b577.svg" alt="Repobeats analytics image" />
</p>

<h2 align="center">Features</h2>

<ul align="center" style="list-style-position: inside; max-width: 600px; margin: auto;">
  <li>Interactive UI for selecting card types and filling required fields</li>
  <li>Supports multiple card types: TaskForce, Management, Site02, MetalCase</li>
  <li>Validates input such as containment levels and permission colors</li>
  <li>Generates the full custom keycard command string</li>
  <li>Copy generated command to clipboard with a simple key press</li>
</ul>

<h2 align="center">Building the Project</h2>

<p align="center">
  You need <a href="https://golang.org/dl/" target="_blank" rel="noopener noreferrer">Go</a> installed.
</p>

<p align="center">
  Build for your current platform:
</p>

<pre style="max-width: 300px; margin: auto;">
<code>go build -o KCard main.go</code>
</pre>

<h2 align="center">Installation</h2>

<p align="center">
  If you don't want to build from source, you can find pre-built versions in the release tab.
</p>
