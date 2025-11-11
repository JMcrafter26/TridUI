Unicode true

####
## Please note: Template replacements don't work in this file. They are provided with default defines like
## mentioned underneath.
## If the keyword is not defined, "wails_tools.nsh" will populate them with the values from ProjectInfo.
## If they are defined here, "wails_tools.nsh" will not touch them. This allows to use this project.nsi manually
## from outside of Wails for debugging and development of the installer.
##
## For development first make a wails nsis build to populate the "wails_tools.nsh":
## > wails build --target windows/amd64 --nsis
## Then you can call makensis on this file with specifying the path to your binary:
## For a AMD64 only installer:
## > makensis -DARG_WAILS_AMD64_BINARY=..\..\bin\app.exe
## For a ARM64 only installer:
## > makensis -DARG_WAILS_ARM64_BINARY=..\..\bin\app.exe
## For a installer with both architectures:
## > makensis -DARG_WAILS_AMD64_BINARY=..\..\bin\app-amd64.exe -DARG_WAILS_ARM64_BINARY=..\..\bin\app-arm64.exe
####
## The following information is taken from the ProjectInfo file, but they can be overwritten here.
####
## !define INFO_PROJECTNAME    "MyProject" # Default "{{.Name}}"
## !define INFO_COMPANYNAME    "MyCompany" # Default "{{.Info.CompanyName}}"
## !define INFO_PRODUCTNAME    "MyProduct" # Default "{{.Info.ProductName}}"
## !define INFO_PRODUCTVERSION "1.0.0"     # Default "{{.Info.ProductVersion}}"
## !define INFO_COPYRIGHT      "Copyright" # Default "{{.Info.Copyright}}"
###
## !define PRODUCT_EXECUTABLE  "Application.exe"      # Default "${INFO_PROJECTNAME}.exe"
## !define UNINST_KEY_NAME     "UninstKeyInRegistry"  # Default "${INFO_COMPANYNAME}${INFO_PRODUCTNAME}"
####
!define REQUEST_EXECUTION_LEVEL "user"            # Set to "user" since we're installing to %APPDATA%
####
## Include the wails tools
####
!include "wails_tools.nsh"

# The version information for this two must consist of 4 parts
VIProductVersion "${INFO_PRODUCTVERSION}.0"
VIFileVersion    "${INFO_PRODUCTVERSION}.0"

VIAddVersionKey "CompanyName"     "${INFO_COMPANYNAME}"
VIAddVersionKey "FileDescription" "${INFO_PRODUCTNAME} Installer"
VIAddVersionKey "ProductVersion"  "${INFO_PRODUCTVERSION}"
VIAddVersionKey "FileVersion"     "${INFO_PRODUCTVERSION}"
VIAddVersionKey "LegalCopyright"  "${INFO_COPYRIGHT}"
VIAddVersionKey "ProductName"     "${INFO_PRODUCTNAME}"
BrandingText "${INFO_PRODUCTNAME} Installer by Cufiy" # This will be shown in the title bar of the installer.

# Enable HiDPI support. https://nsis.sourceforge.io/Reference/ManifestDPIAware
ManifestDPIAware true

!include "MUI.nsh"
!include "nsDialogs.nsh"
!include "LogicLib.nsh"

!define MUI_ICON "..\icon.ico"
!define MUI_UNICON "..\icon.ico"
# !define MUI_WELCOMEFINISHPAGE_BITMAP "resources\leftimage.bmp" #Include this to add a bitmap on the left side of the Welcome Page. Must be a size of 164x314
!define MUI_FINISHPAGE_NOAUTOCLOSE # Wait on the INSTFILES page so the user can take a look into the details of the installation steps
!define MUI_FINISHPAGE_RUN "$INSTDIR\${PRODUCT_EXECUTABLE}" # Run the application after installation
!define MUI_FINISHPAGE_RUN_TEXT_VAR LaunchAppText
!define MUI_FINISHPAGE_SHOWREADME ""
!define MUI_FINISHPAGE_SHOWREADME_NOTCHECKED
!define MUI_FINISHPAGE_SHOWREADME_TEXT_VAR CreateDesktopShortcutTextVar
!define MUI_FINISHPAGE_SHOWREADME_FUNCTION CreateDesktopShortcut
!define MUI_ABORTWARNING # This will warn the user if they exit from the installer.

!insertmacro MUI_PAGE_WELCOME # Welcome to the installer page.
# !insertmacro MUI_PAGE_LICENSE "resources\eula.txt" # Adds a EULA page to the installer
!insertmacro MUI_PAGE_DIRECTORY # In which folder install page.
!insertmacro MUI_PAGE_INSTFILES # Installing page.
!insertmacro MUI_PAGE_FINISH # Finished installation page.


!define MUI_UNPAGE_WELCOME
!insertmacro MUI_UNPAGE_WELCOME
UninstPage custom un.SorryPageCreate un.SorryPageLeave
!insertmacro MUI_UNPAGE_INSTFILES # Uninstalling page

# Multilingual support - NSIS will automatically select based on system language
!insertmacro MUI_LANGUAGE "English"
!insertmacro MUI_LANGUAGE "German"
!insertmacro MUI_LANGUAGE "Spanish"
!insertmacro MUI_LANGUAGE "French"
!insertmacro MUI_LANGUAGE "Italian"
!insertmacro MUI_LANGUAGE "Japanese"
!insertmacro MUI_LANGUAGE "Polish"
!insertmacro MUI_LANGUAGE "Portuguese"
!insertmacro MUI_LANGUAGE "Russian"
!insertmacro MUI_LANGUAGE "SimpChinese"

# Include translation files
!include "translations\English.nsh"
!include "translations\German.nsh"
!include "translations\Spanish.nsh"
!include "translations\French.nsh"
!include "translations\Italian.nsh"
!include "translations\Japanese.nsh"
!include "translations\Polish.nsh"
!include "translations\Portuguese.nsh"
!include "translations\Russian.nsh"
!include "translations\SimpChinese.nsh"

## The following two statements can be used to sign the installer and the uninstaller. The path to the binaries are provided in %1
#!uninstfinalize 'signtool --file "%1"'
#!finalize 'signtool --file "%1"'

Name "${INFO_PRODUCTNAME}"
OutFile "..\..\bin\${INFO_PROJECTNAME}-${ARCH}-installer.exe" # Name of the installer's file.
InstallDir "$AppData\TridUI" # Default installing folder is %APPDATA%\TridUI
ShowInstDetails show # This will always show the installation details.

Var LaunchAppText
Var CreateDesktopShortcutTextVar

Function .onInit
   !insertmacro wails.checkArchitecture
   
   ; Set multilingual strings based on selected language
    StrCpy $LaunchAppText $(LaunchApp)
    StrCpy $CreateDesktopShortcutTextVar $(CreateDesktopShortcutText)
FunctionEnd

Function CreateDesktopShortcut
    CreateShortCut "$DESKTOP\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
FunctionEnd

Section
    !insertmacro wails.setShellContext

    !insertmacro wails.webview2runtime

    SetOutPath $INSTDIR

    !insertmacro wails.files

    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

    !insertmacro wails.associateFiles
    !insertmacro wails.associateCustomProtocols

    !insertmacro wails.writeUninstaller
    
    ; Ensure the uninstaller is registered in Windows Add/Remove Programs
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "DisplayName" "${INFO_PRODUCTNAME}"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "UninstallString" "$INSTDIR\uninstall.exe"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "InstallLocation" "$INSTDIR"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "Publisher" "${INFO_COMPANYNAME}"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "DisplayVersion" "${INFO_PRODUCTVERSION}"
    WriteRegDWORD HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "NoModify" 1
    WriteRegDWORD HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}" "NoRepair" 1
SectionEnd

!define MUI_UNWELCOMEPAGE_TITLE_VAR UninstallTitleText
!define MUI_UNWELCOMEPAGE_TEXT_VAR UninstallWelcomeText

Var UninstallTitleText
Var UninstallWelcomeText

Function un.onInit
    !insertmacro MUI_UNGETLANGUAGE
    
    ; Set multilingual strings based on selected language
    StrCpy $UninstallTitleText $(UninstallTitle)
    StrCpy $UninstallWelcomeText $(UninstallWelcome)
FunctionEnd

; Variables for uninstaller
Var SorryDialog
Var GitHubButton

Function un.SorryPageCreate
    !insertmacro MUI_HEADER_TEXT $(SorryTitle) $(SorrySubtitle)
    
    nsDialogs::Create 1018
    Pop $SorryDialog
    
    ${If} $SorryDialog == error
        Abort
    ${EndIf}
    
    ${NSD_CreateLabel} 10u 10u 90% 25u $(SorryLine1)
    Pop $0

    ${NSD_CreateLabel} 10u 35u 90% 35u $(SorryLine2)
    Pop $0
    
    ${NSD_CreateLabel} 20u 70u 70% 15u $(SorryLine3)
    Pop $0

    ${NSD_CreateButton} 10u 90u 120u 20u $(OpenGitHubButton)
    Pop $GitHubButton
    ${NSD_OnClick} $GitHubButton un.OpenGitHub
    
    ${NSD_CreateLabel} 10u 120u 90% 15u $(ThankYou)
    Pop $0
    
    nsDialogs::Show
FunctionEnd

Function un.SorryPageLeave
    ; Nothing to do here
FunctionEnd

Function un.OpenGitHub
    ExecShell "open" "https://github.com/JMcrafter26/TridUI"
FunctionEnd

Section "uninstall"
    !insertmacro wails.setShellContext

    ; Remove entire installation directory and all contents
    RMDir /r $INSTDIR
    
    ; Remove WebView2 DataPath
    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}"

    ; Remove triddefs.trd and triddefs.trd.meta
    Delete "$AppData\TridUI\triddefs.trd"
    Delete "$AppData\TridUI\triddefs.trd.meta"
    
    ; Remove application data folder in AppData\Roaming
    RMDir /r "$AppData\TridUI"

    ; Always remove shortcuts and registry entries
    Delete "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk"
    Delete "$DESKTOP\${INFO_PRODUCTNAME}.lnk"

    ; Remove from Add/Remove Programs
    DeleteRegKey HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}"

    !insertmacro wails.unassociateFiles
    !insertmacro wails.unassociateCustomProtocols

    !insertmacro wails.deleteUninstaller
SectionEnd
