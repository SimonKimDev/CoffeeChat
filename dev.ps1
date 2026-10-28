[CmdletBinding()]
param(
    [Parameter(Position = 0)]
    [ValidateSet('help', 'build', 'run', 'test', 'vet', 'check', 'tidy', 'clean')]
    [string]$Task = 'help'
)

$ErrorActionPreference = 'Stop'

function Invoke-Go {
    param([string[]]$GoArguments)

    & go @GoArguments
    if ($LASTEXITCODE -ne 0) {
        exit $LASTEXITCODE
    }
}

if ($Task -eq 'help') {
    Write-Output @'
Usage: .\dev.ps1 [task]

Tasks:
  help   Show this help (default)
  build  Build the relay into ./bin
  run    Build and run the relay
  test   Run all Go tests
  vet    Run Go's static checks
  check  Build, test, and vet
  tidy   Tidy Go module dependencies
  clean  Remove build output
'@
    return
}

$binaryName = if ($env:OS -eq 'Windows_NT') { 'coffeechat-server.exe' } else { 'coffeechat-server' }
$binaryPath = Join-Path $PSScriptRoot "bin\$binaryName"

Push-Location $PSScriptRoot
try {
    switch ($Task) {
        { $_ -in 'build', 'run', 'check' } {
            Invoke-Go -GoArguments @('build', '-o', $binaryPath, './cmd/server')
            if ($Task -eq 'run') {
                & $binaryPath
                if ($LASTEXITCODE -ne 0) {
                    exit $LASTEXITCODE
                }
            }
            if ($Task -eq 'check') {
                Invoke-Go -GoArguments @('test', './...')
                Invoke-Go -GoArguments @('vet', './...')
            }
        }
        'test' { Invoke-Go -GoArguments @('test', './...') }
        'vet' { Invoke-Go -GoArguments @('vet', './...') }
        'tidy' { Invoke-Go -GoArguments @('mod', 'tidy') }
        'clean' {
            $outputPath = Join-Path $PSScriptRoot 'bin'
            if (Test-Path -LiteralPath $outputPath) {
                Remove-Item -LiteralPath $outputPath -Recurse -Force
            }
        }
    }
}
finally {
    Pop-Location
}
