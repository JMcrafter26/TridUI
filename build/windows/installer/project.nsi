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
!define MUI_FINISHPAGE_RUN_TEXT "Launch ${INFO_PRODUCTNAME}"
!define MUI_ABORTWARNING # This will warn the user if they exit from the installer.

!insertmacro MUI_PAGE_WELCOME # Welcome to the installer page.
# !insertmacro MUI_PAGE_LICENSE "resources\eula.txt" # Adds a EULA page to the installer
!insertmacro MUI_PAGE_DIRECTORY # In which folder install page.
!insertmacro MUI_PAGE_INSTFILES # Installing page.
!insertmacro MUI_PAGE_FINISH # Finished installation page.


!define MUI_UNPAGE_WELCOME
!insertmacro MUI_UNPAGE_WELCOME
UninstPage custom un.SorryPageCreate un.SorryPageLeave
UninstPage custom un.ComponentsPageCreate un.ComponentsPageLeave
!insertmacro MUI_UNPAGE_INSTFILES # Uninstalling page

!insertmacro MUI_LANGUAGE "English" # Set the Language of the installer

## The following two statements can be used to sign the installer and the uninstaller. The path to the binaries are provided in %1
#!uninstfinalize 'signtool --file "%1"'
#!finalize 'signtool --file "%1"'

Name "${INFO_PRODUCTNAME}"
OutFile "..\..\bin\${INFO_PROJECTNAME}-${ARCH}-installer.exe" # Name of the installer's file.
InstallDir "$AppData\TridscanUI" # Default installing folder is %APPDATA%\TridscanUI
ShowInstDetails show # This will always show the installation details.

Function .onInit
   !insertmacro wails.checkArchitecture
FunctionEnd

Section
    !insertmacro wails.setShellContext

    !insertmacro wails.webview2runtime

    SetOutPath $INSTDIR

    !insertmacro wails.files

    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    CreateShortCut "$DESKTOP\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

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

Function un.onInit
    !insertmacro MUI_UNGETLANGUAGE
FunctionEnd

!define MUI_UNWELCOMEPAGE_TITLE "Uninstalling TridscanUI"
!define MUI_UNWELCOMEPAGE_TEXT "Welcome to the TridscanUI uninstaller."

; Variables for uninstaller
Var RemoveAllData
Var Dialog
Var RadioButton1
Var RadioButton2
Var SorryDialog
Var GitHubButton

Function un.SorryPageCreate
    !insertmacro MUI_HEADER_TEXT "We're Sorry to See You Go" "Thank you for trying TridscanUI"
    
    nsDialogs::Create 1018
    Pop $SorryDialog
    
    ${If} $SorryDialog == error
        Abort
    ${EndIf}
    
    ${NSD_CreateLabel} 10u 10u 90% 25u "We're sorry to see you go."
    Pop $0

    ${NSD_CreateLabel} 10u 35u 90% 35u "If you ran into any issues, have suggestions, or just want to share your experience, please let us know. You can always reach out or get help using GitHub:"
    Pop $0
    
    ${NSD_CreateLabel} 20u 70u 70% 15u "- Need assistance with TridscanUI? Have feedback/ideas? Want to stay in touch?"
    Pop $0

    ${NSD_CreateButton} 10u 90u 120u 20u "Open GitHub page"
    Pop $GitHubButton
    ${NSD_OnClick} $GitHubButton un.OpenGitHub
    
    ${NSD_CreateLabel} 10u 120u 90% 15u "Thank you for trying TridscanUI!"
    Pop $0
    
    nsDialogs::Show
FunctionEnd

Function un.SorryPageLeave
    ; Nothing to do here
FunctionEnd

Function un.OpenGitHub
    ExecShell "open" "https://github.com/JMcrafter26/TridscanUI"
FunctionEnd

; Uninstaller sections with different options
Section /o "un.Remove all data (including downloaded content)" un.RemoveAllData
SectionEnd

Section "un.Remove program only (keep downloaded content)" un.RemoveProgram
SectionEnd

Function un.ComponentsPageCreate
    !insertmacro MUI_HEADER_TEXT "Choose Uninstall Type" "Select what you want to remove"
    
    nsDialogs::Create 1018
    Pop $Dialog
    
    ${If} $Dialog == error
        Abort
    ${EndIf}
    
    ${NSD_CreateLabel} 0 10u 100% 20u "How would you like to uninstall TridscanUI?"
    Pop $0
    
    ${NSD_CreateRadioButton} 10u 40u 90% 15u "Remove program only (keep downloaded content)"
    Pop $RadioButton1
    ${NSD_Check} $RadioButton1
    
    ${NSD_CreateRadioButton} 10u 60u 90% 15u "Complete removal (remove all data including downloads)"
    Pop $RadioButton2
    
    ${NSD_CreateLabel} 0 85u 100% 40u "Note: If you choose complete removal, all your downloaded content and settings will be permanently deleted. If you only want to remove the program, your downloaded content will be preserved."
    Pop $0
    
    nsDialogs::Show
FunctionEnd

Function un.ComponentsPageLeave
    ${NSD_GetState} $RadioButton2 $RemoveAllData
FunctionEnd

Section "uninstall"
    !insertmacro wails.setShellContext

    ; Always remove the main executable and core files
    Delete "$INSTDIR\${PRODUCT_EXECUTABLE}"
    
    ; Remove sandbox files if they exist
    IfFileExists "$INSTDIR\sandbox_host.exe" 0 +2
        Delete "$INSTDIR\sandbox_host.exe"
    IfFileExists "$INSTDIR\sandbox_host.zip" 0 +2
        Delete "$INSTDIR\sandbox_host.zip"
    
    ; Remove WebView2 DataPath
    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}"
    
    ; Check if user wants complete removal
    ${If} $RemoveAllData == 1
        ; Complete removal - remove entire directory
        RMDir /r $INSTDIR
    ${Else}
        ; Partial removal - only remove specific files, keep user data
        ; Remove other application files but preserve user content
        Delete "$INSTDIR\*.dll"
        Delete "$INSTDIR\*.json"
        Delete "$INSTDIR\*.log"
        ; Keep the directory structure and any user content
    ${EndIf}

    ; Always remove shortcuts and registry entries
    Delete "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk"
    Delete "$DESKTOP\${INFO_PRODUCTNAME}.lnk"

    ; Remove from Add/Remove Programs
    DeleteRegKey HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\${INFO_PRODUCTNAME}"

    !insertmacro wails.unassociateFiles
    !insertmacro wails.unassociateCustomProtocols

    !insertmacro wails.deleteUninstaller
SectionEnd
