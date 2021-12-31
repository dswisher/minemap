// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using CommandLine;
using MineMap.Cli.Commands;
using MineMap.Cli.Exceptions;
using MineMap.Cli.Options;

namespace MineMap.Cli
{
    public static class Program
    {
        public static int Main(string[] args)
        {
            try
            {
                var parsedArgs = Parser.Default.ParseArguments<DumpChunkOptions, DumpRegionOptions, TimeMapOptions>(args);

                // Process each verb separately, because that's the way the library seems to work...is there a better way?
                parsedArgs.WithParsed<DumpChunkOptions>(options => new DumpChunkCommand().Run(options));
                parsedArgs.WithParsed<DumpRegionOptions>(options => new DumpRegionCommand().Run(options));
                parsedArgs.WithParsed<TimeMapOptions>(options => new TimeMapCommand().Run(options));
            }
            catch (CliException ex)
            {
                Console.WriteLine(ex.Message);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Unhandled exception in main.");
                Console.WriteLine(ex);
            }

            return 0;
        }
    }
}
