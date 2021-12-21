// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using CommandLine;
using MineMap.Cli.Commands;
using MineMap.Cli.Options;

namespace MineMap.Cli
{
    public static class Program
    {
        public static int Main(string[] args)
        {
            try
            {
                // TODO - remove "object" once we have a second verb
                var parsedArgs = Parser.Default.ParseArguments<DumpRegionOptions, object>(args);

                // Process each verb separately, because that's the way the library seems to work...is there a better way?
                parsedArgs.WithParsed<DumpRegionOptions>(options => new DumpRegionCommand().Run(options));
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
