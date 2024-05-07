function Start-InstallWindowsNextGenNetworkingIfNeeded {
    param(
        [Parameter(Mandatory = $true)]
        [bool]$EnableInstall,
        [Parameter(Mandatory = $true)]
        [string]$UrlBase
    )

    if (-not $EnableInstall) {
        Write-Log "EnableNextGenNetworking is false, so skip installation of Windows eBPF"
        return
    }

    Logs-To-Event -TaskName "AKS.WindowsCSE.InstallWindowsNextGenNetworking" -TaskMessage "Start to install Windows next-gen networking eBPF using URL base $UrlBase"

    # TODO: this is a placeholder for the installation script.
}
