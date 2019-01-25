// Code generated by go generate; DO NOT EDIT.
package main

const infoPlistTmpl = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>{{.DevRegion}}</string>

	<key>CFBundleExecutable</key>
	<string>{{.Executable}}</string>

	<key>CFBundleIconFile</key>
	<string>{{.Icon}}</string>

	<key>CFBundleIdentifier</key>
	<string>{{.ID}}</string>

	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>

	<key>CFBundleName</key>
	<string>{{.Name}}</string>

	<key>CFBundlePackageType</key>
	<string>APPL</string>

	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>MacOSX</string>
	</array>

	<key>CFBundleShortVersionString</key>
	<string>{{.Version}}</string>

	<key>CFBundleVersion</key>
	<string>{{.BuildNumber}}</string>

	<key>LSMinimumSystemVersion</key>
	<string>{{.DeploymentTarget}}</string>

	<key>LSApplicationCategoryType</key>
	<string>{{.Category}}</string>

	{{if .Background}}
	<key>LSUIElement</key>
	<true/>
	{{end}}

	<key>NSHumanReadableCopyright</key>
	<string>{{html .Copyright}}</string>

	<key>NSPrincipalClass</key>
	<string>NSApplication</string>

	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSAllowsArbitraryLoadsInWebContent</key>
		<true/>
	</dict>

	<key>CFBundleDocumentTypes</key>
	<array>
		{{range .SupportedFiles}}
		<dict>
			<key>CFBundleTypeName</key>
			<string>{{.Name}}</string>
			<key>CFBundleTypeRole</key>
			<string>{{.Role}}</string>
			{{if .Icon}}
			<key>CFBundleTypeIconFile</key>
			<string>{{.Icon}}</string>
			{{end}}
			<key>LSItemContentTypes</key>
			<array>{{range .UTIs}}
				<string>{{.}}</string>{{end}}
			</array>
		</dict>
		{{end}}
	</array>

	<key>CFBundleURLTypes</key>
	<array>
		<dict>
			<key>CFBundleURLName</key>
			<string>{{.ID}}</string>
			<key>CFBundleTypeRole</key>
			<string>Editor</string>
			<key>CFBundleURLSchemes</key>
			<array>
				<string>{{.URLScheme}}</string>
			</array>
		</dict>
	</array>
</dict>
</plist>`

const entitlementsPlistTmpl = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	{{if .Sandbox}}
	<key>com.apple.security.app-sandbox</key>
	<true/>
	{{end}}

	<!-- Network -->
	{{if .Server}}
	<key>com.apple.security.network.server</key>
	<true/>
	{{end}}
	<key>com.apple.security.network.client</key>
	<true/>

	<!-- Hadrware -->
	{{if .Camera}}
	<key>com.apple.security.device.camera</key>
	<true/>
	{{end}}
	{{if .Microphone}}
	<key>com.apple.security.device.microphone</key>
	<true/>
	{{end}}
	{{if .USB}}
	<key>com.apple.security.device.usb</key>
	<true/>
	{{end}}
	{{if .Printers}}
	<key>com.apple.security.print</key>
	<true/>
	{{end}}
	{{if .Bluetooth}}
	<key>com.apple.security.device.bluetooth</key>
	<true/>
	{{end}}

	<!-- AppData -->
	{{if .Contacts}}
	<key>com.apple.security.personal-information.addressbook</key>
	<true/>
	{{end}}
	{{if .Location}}
	<key>com.apple.security.personal-information.location</key>
	<true/>
	{{end}}
	{{if .Calendar}}
	<key>com.apple.security.personal-information.calendars</key>
	<true/>
	{{end}}

	<!-- FileAccess -->
	{{if len .FilePickers}}
	<key>com.apple.security.files.user-selected.{{.FilePickers}}</key>
	<true/>
	{{end}}
	{{if len .Downloads}}
	<key>com.apple.security.files.downloads.{{.Downloads}}</key>
	<true/>
	{{end}}
	{{if len .Pictures}}
	<key>com.apple.security.assets.pictures.{{.Pictures}}/key>
	<true/>
	{{end}}
	{{if len .Music}}
	<key>com.apple.security.assets.music.{{.Music}}</key>
	<true/>
	{{end}}
	{{if len .Movies}}
	<key>com.apple.security.assets.movies.{{.Movies}}/key>
	<true/>
	{{end}}
</dict>
</plist>`

const appxManifestTmpl = `
<?xml version="1.0" encoding="utf-8"?>
<Package xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10" 
    xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" 
    xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities" 
    xmlns:desktop="http://schemas.microsoft.com/appx/manifest/desktop/windows10" IgnorableNamespaces="uap rescap desktop">
    <Identity Name="{{.ID}}" Publisher="CN={{.Publisher}}" Version="1.0.0.0" />
    <Properties>
        <DisplayName>{{.Name}}</DisplayName>
        <PublisherDisplayName>{{.Publisher}}</PublisherDisplayName>
        <Logo>Assets\StoreLogo.png</Logo>
    </Properties>
    <Dependencies>
        <TargetDeviceFamily Name="Windows.Desktop" MinVersion="10.0.15063.0" MaxVersionTested="10.0.17134.0" />
        <PackageDependency Name="Microsoft.NET.Native.Framework.1.7" MinVersion="1.7.25531.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
        <PackageDependency Name="Microsoft.NET.Native.Runtime.1.7" MinVersion="1.7.25531.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
        <PackageDependency Name="Microsoft.VCLibs.140.00" MinVersion="14.0.22929.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
    </Dependencies>
    <Resources>
        <Resource Language="EN-US" />
        <Resource uap:Scale="100" />
        <Resource uap:Scale="125" />
        <Resource uap:Scale="150" />
        <Resource uap:Scale="200" />
        <Resource uap:Scale="400" />
    </Resources>
    <Applications>
        <Application Id="{{.ID}}" Executable="uwp.exe" EntryPoint="uwp.App">
            <uap:VisualElements DisplayName="{{.Name}}" Description="{{.Description}}" BackgroundColor="transparent" Square150x150Logo="Assets\Square150x150Logo.png" Square44x44Logo="Assets\Square44x44Logo.png">
                <uap:DefaultTile Wide310x150Logo="Assets\Wide310x150Logo.png" Square310x310Logo="Assets\Square310x310Logo.png" Square71x71Logo="Assets\Square71x71Logo.png">
                    <uap:ShowNameOnTiles>
                        <uap:ShowOn Tile="square150x150Logo" />
                        <uap:ShowOn Tile="wide310x150Logo" />
                        <uap:ShowOn Tile="square310x310Logo" />
                    </uap:ShowNameOnTiles>
                </uap:DefaultTile>
            </uap:VisualElements>
            <Extensions>
                <desktop:Extension Category="windows.fullTrustProcess" Executable="{{.Executable}}" />
                <uap:Extension Category="windows.appService">
                    <uap:AppService Name="InProcessAppService" />
                </uap:Extension>
                <uap:Extension Category="windows.protocol">
                    <uap:Protocol Name="{{.URLScheme}}" />
                </uap:Extension>
                {{range .SupportedFiles}}
                <uap:Extension Category="windows.fileTypeAssociation">
                    <uap:FileTypeAssociation Name="{{.Name}}">
                        <uap:InfoTip>{{.Help}}</uap:InfoTip>
                        {{if .Icon}}
                        <uap:Logo>{{.Icon}}</uap:Logo>
                        {{end}}
                        {{range .Extensions}}
                        <uap:SupportedFileTypes>
                            {{if .Mime}}
                            <uap:FileType ContentType="{{.Mime}}">{{.Ext}}</uap:FileType>
                            {{else}}
                            <uap:FileType>{{.Ext}}</uap:FileType>
                            {{end}}
                        </uap:SupportedFileTypes>
                        {{end}}
                    </uap:FileTypeAssociation>
                </uap:Extension>
                {{end}}
            </Extensions>
        </Application>
    </Applications>
    <Capabilities>
        <Capability Name="internetClient" />
        <rescap:Capability Name="runFullTrust"/>
        <rescap:Capability Name="confirmAppClose"/>
    </Capabilities>
</Package>`
