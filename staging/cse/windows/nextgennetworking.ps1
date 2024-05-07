function Start-InstallWindowsNextGenNetworkingIfNeeded {
    param(
        [Parameter(Mandatory = $true)]
        [bool]$EnableInstall,
        [Parameter(Mandatory = $true)]
        [AllowEmptyString()]
        [string]$URLBase
    )

    if (-not $EnableInstall) {
        Write-Log "EnableNextGenNetworking is false, so skip installation of Windows eBPF networking"
        return
    }

    Logs-To-Event -TaskName "AKS.WindowsCSE.InstallWindowsNextGenNetworking" -TaskMessage "Start to install Windows next-gen networking eBPF using URL base $URLBase"

    # TODO: this is a placeholder for the installation script.
}
