function Start-InstallWindowsNextGenNetworking {
    param(
        [Parameter(Mandatory = $true)]
        [string]$UrlBase
    )

    Logs-To-Event -TaskName "AKS.WindowsCSE.InstallWindowsNextGenNetworking" -TaskMessage "Start to install Windows next-gen networking eBPF using URL base $UrlBase"

    # TODO: install Windows eBPF
}
